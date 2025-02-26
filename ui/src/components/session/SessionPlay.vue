<template>
  <fragment
    v-if="hidden()"
    data-test="render-fragment"
  >
    <v-tooltip
      v-if="recorded"
      bottom
    >
      <template #activator="{ on }">
        <span v-on="on">
          <v-icon
            :disabled="!isOwner"
            v-on="on"
            @click="displayDialog"
          >
            mdi-play-circle
          </v-icon>
        </span>
      </template>

      <div>
        <span
          v-if="isOwner"
          data-test="text-tooltip"
        >
          Play
        </span>

        <span v-else>
          You are not the owner of this namespace
        </span>
      </div>
    </v-tooltip>

    <v-dialog
      v-model="dialog"
      :max-width="1024"
      :transition="false"
    >
      <v-card
        :elevation="0"
        data-test="sessionPlay-card"
      >
        <v-toolbar
          dark
          color="primary"
        >
          <v-btn
            icon
            dark
            @click="dialog = !dialog"
          >
            <v-icon>close</v-icon>
          </v-btn>
          <v-toolbar-title>Watch Session</v-toolbar-title>
          <v-spacer />
        </v-toolbar>
        <div ref="playterminal" />
        <v-container class="pa-0">
          <v-row no-gutters>
            <v-col
              cols="2"
              sm="6"
              md="1"
            >
              <v-card
                :elevation="0"
                class="pt-4 ml-7"
                tile
              >
                <v-icon
                  v-if="!paused"
                  large
                  class="pl-0"
                  color="primary"
                  @click="pauseHandler"
                >
                  mdi-pause-circle
                </v-icon>
                <v-icon
                  v-else
                  large
                  class="pl-0"
                  color="primary"
                  @click="pauseHandler"
                >
                  mdi-play-circle
                </v-icon>
              </v-card>
            </v-col>
            <v-col
              cols="6"
              md="9"
            >
              <v-card
                :elevation="0"
                class="pt-4 pl-9 mr-5"
                tile
              >
                <v-slider
                  v-model="currentTime"
                  class="ml-0"
                  min="0"
                  :max="totalLength"
                  :label="`${nowTimerDisplay} - ${endTimerDisplay}`"
                  @change="changeSliderTime"
                  @mousedown="previousPause=paused, paused=true"
                  @mouseup="paused=previousPause"
                />
              </v-card>
            </v-col>
            <v-col
              cols="6"
              md="2"
            >
              <v-card
                :elevation="0"
                class="pt-4 ml-5"
                tile
              >
                <v-select
                  v-model="defaultSpeed"
                  class="pr-8 mt-0 pt-0 mr-4"
                  :items="speedList"
                  menu-props="auto"
                  prepend-icon="mdi-speedometer"
                  @change="speedChange"
                >
                  <template #selection="{ item }">
                    <span> {{ item }} </span>
                  </template>
                </v-select>
              </v-card>
            </v-col>
          </v-row>
        </v-container>
      </v-card>
    </v-dialog>
  </fragment>
</template>

<script>

import { Terminal } from 'xterm';
import { FitAddon } from 'xterm-addon-fit';
import moment from 'moment';
import 'moment-duration-format';
import 'xterm/css/xterm.css';

export default {
  name: 'SessionPlay',

  props: {
    uid: {
      type: String,
      required: true,
    },
    recorded: {
      type: Boolean,
      required: true,
    },
  },

  data() {
    return {
      dialog: false,
      currentTime: 0,
      totalLength: 0,
      endTimerDisplay: 0,
      getTimerNow: 0,
      paused: false,
      previousPause: false,
      sliderChange: false,
      speedList: [0.5, 1, 1.5, 2, 4],
      logs: [],
      frames: [],
      defaultSpeed: 1,
      transition: false,
    };
  },

  computed: {
    length() {
      return this.logs.length;
    },

    nowTimerDisplay() {
      return this.getTimerNow;
    },

    isOwner() {
      return this.$store.getters['namespaces/owner'];
    },
  },

  watch: {
    dialog(value) {
      if (!value) {
        this.close();
      }
    },
  },

  updated() {
    if (this.dialog) {
      this.getTimerNow = this.getDisplaySliderInfo(this.currentTime).display;
    }
  },

  methods: {
    async openPlay() {
      if (this.recorded) {
        // receive data
        await this.$store.dispatch('sessions/getLogSession', this.uid);
        this.logs = this.$store.getters['sessions/get'];
        this.totalLength = this.getDisplaySliderInfo(null).intervalLength;
        this.endTimerDisplay = this.getDisplaySliderInfo(null).display;
        this.getTimerNow = this.getDisplaySliderInfo(this.currentTime).display;
        this.frames = this.createFrames();

        this.xterm = new Terminal({ // instantiate Terminal
          cursorBlink: true,
          fontFamily: 'monospace',
        });

        this.fitAddon = new FitAddon(); // load fit
        this.xterm.loadAddon(this.fitAddon); // adjust screen in container
        if (this.xterm.element) {
          this.xterm.reset();
        }
      }
    },

    async displayDialog() { // await to change dialog for the connection
      try {
        await this.openPlay();
        this.dialog = !this.dialog;
        this.$nextTick().then(() => {
          this.connect();
        });
      } catch {
        this.$store.dispatch('snackbar/showSnackbarErrorLoading', this.$errors.snackbar.sessionPlay);
      }
    },

    async connect() {
      this.xterm.open(this.$refs.playterminal);
      this.fitAddon.fit();
      this.xterm.focus();
      this.print(0, this.logs);
      this.timer();
    },

    getDisplaySliderInfo(timeMs) {
      let interval;

      if (!timeMs) { // not params, will return metrics to max timelengtht
        const max = new Date(this.logs[this.length - 1].time);
        const min = new Date(this.logs[0].time);
        interval = max - min;
      } else { // it will format to the time argument passed
        interval = timeMs;
      }

      const duration = moment.duration(interval, 'milliseconds');

      // format according to how long
      let hoursFormat;
      if (duration.asHours() > 1) hoursFormat = 'h';
      else hoursFormat = '';

      const displayTime = duration.format(`${hoursFormat}:mm:ss`, {
        trim: false,
      });

      return {
        display: displayTime, // format the slider label
        intervalLength: interval, // length of slider in Ms
      };
    },

    createFrames() { // create cumulative frames for the exibition in slider
      let time = 0;
      let message = '';
      const arrFrames = [{
        incMessage: message += this.logs[0].message,
        incTime: time,
      }];

      for (let i = 1; i < this.logs.length; i += 1) {
        const future = new Date(this.logs[i].time);
        const now = new Date(this.logs[i - 1].time);
        const interval = moment.duration(future - now, 'milliseconds').asMilliseconds();
        time += interval;
        message += this.logs[i].message;
        arrFrames.push({
          incMessage: message,
          incTime: time,
        });
      }
      return arrFrames;
    },

    speedChange(speed) {
      this.defaultSpeed = speed;
      this.xtermSyncFrame(this.currentTime);
    },

    timer() { // Increments the slider
      if (!this.paused) {
        if (this.currentTime >= this.totalLength) return;
        this.currentTime += 100;
      }
      this.iterativeTimer = setTimeout(this.timer.bind(null), 100 * (1 / this.defaultSpeed));
    },

    changeSliderTime() { // Moving the Slider
      this.sliderChange = true;
      this.xtermSyncFrame(this.currentTime);
    },

    pauseHandler() {
      this.paused = !this.paused;
      this.xtermSyncFrame(this.currentTime);
    },

    close() {
      this.transition = true;
      if (this.xterm) this.xterm.dispose();
      this.clear();
      this.currentTime = 0;
      this.paused = false;
      this.defaultSpeed = 1;
    },

    clear() { // Ensure to clear functions for syncronism
      clearInterval(this.iterativePrinting);
      clearInterval(this.iterativeTimer);
    },

    xtermSyncFrame(givenTime) {
      this.xterm.write('\u001Bc'); // clean screen
      const frame = this.searchClosestFrame(givenTime, this.frames);
      this.clear();
      this.xterm.write(frame.message); // write frame on xterm
      this.iterativeTimer = setTimeout(this.timer.bind(null),
        1);
      this.iterativePrinting = setTimeout(this.print.bind(null, frame.index + 1, this.logs),
        frame.waitForPrint * (1 / this.defaultSpeed));
    },

    searchClosestFrame(givenTime, frames) { // applies a binary search to find nearest frame
      let between;
      let lowerBound = 0;
      let higherBound = frames.length - 1;
      let nextTimeSetPrint;

      for (;higherBound - lowerBound > 1;) { // progressive increment search
        between = Math.floor((lowerBound + higherBound) / 2);
        if (frames[between].incTime < givenTime) {
          lowerBound = between;
          nextTimeSetPrint = givenTime - frames[between].incTime;
        } else {
          higherBound = between;
          nextTimeSetPrint = frames[between].incTime - givenTime;
        }
      }
      return {
        message: frames[lowerBound].incMessage,
        index: lowerBound,
        waitForPrint: nextTimeSetPrint,
      };
    },

    print(i, logsArray) { // Writes iteratevely on xterm as time progresses
      this.sliderChange = false;
      if (!this.paused) {
        this.xterm.write(`${logsArray[i].message}`);
        if (i === logsArray.length - 1) return;
        const nowTimerDisplay = new Date(logsArray[i].time);
        const future = new Date(logsArray[i + 1].time);
        const interval = future - nowTimerDisplay;
        this.iterativePrinting = setTimeout(this.print.bind(null, i + 1, logsArray),
          interval * (1 / this.defaultSpeed));
      }
    },

    hidden() {
      return this.$env.isEnterprise;
    },
  },
};

</script>
