<template>
  <fragment>
    <v-tooltip
      v-if="createRule"
      bottom
      :disabled="isOwner"
    >
      <template #activator="{ on }">
        <div v-on="on">
          <v-btn
            :disabled="!isOwner"
            class="v-btn--active"
            text
            color="primary"
            data-test="add-btn"
            @click="dialog = !dialog"
          >
            Add Rule
          </v-btn>
        </div>
      </template>

      <span>
        You are not the owner of this namespace
      </span>
    </v-tooltip>

    <v-tooltip
      v-else
      bottom
    >
      <template #activator="{ on }">
        <span v-on="on">
          <v-icon
            :disabled="!isOwner"
            v-on="on"
            @click="dialog = !dialog"
          >
            edit
          </v-icon>
        </span>
      </template>

      <div>
        <span
          v-if="isOwner"
          data-test="text-tooltip"
        >
          Edit
        </span>

        <span v-else>
          You are not the owner of this namespace
        </span>
      </div>
    </v-tooltip>

    <v-dialog
      v-model="dialog"
      max-width="400"
      @click:outside="close"
    >
      <v-card data-test="firewallRuleForm-card">
        <ValidationObserver
          ref="obs"
          v-slot="{ passes }"
        >
          <v-card-title
            v-if="createRule"
            class="headline grey lighten-2 text-center"
          >
            New Rule
          </v-card-title>
          <v-card-title
            v-else
            class="headline grey lighten-2 text-center"
          >
            Edit Rule
          </v-card-title>

          <v-card-text>
            <v-layout
              justify-space-between
              align-center
            >
              <v-flex>
                <v-card :elevation="0">
                  <v-card-text class="v-label theme--light pl-0">
                    Active
                  </v-card-text>
                </v-card>
              </v-flex>

              <v-flex
                xs2
              >
                <v-card
                  :elevation="0"
                >
                  <v-switch
                    v-model="ruleFirewallLocal.active"
                  />
                </v-card>
              </v-flex>
            </v-layout>
            <ValidationProvider
              v-slot="{ errors }"
              name="Priority"
              rules="required|integer"
            >
              <v-text-field
                v-model="ruleFirewallLocal.priority"
                label="Priority"
                type="number"
                :error-messages="errors"
                required
              />
            </ValidationProvider>

            <ValidationProvider
              v-slot="{ errors }"
              name="Action"
              rules="required"
            >
              <v-select
                v-model="ruleFirewallLocal.action"
                :items="state"
                item-text="name"
                item-value="id"
                label="Action"
                :error-messages="errors"
                required
              />
            </ValidationProvider>

            <ValidationProvider
              v-slot="{ errors }"
              name="Source IP"
              rules="required"
            >
              <v-text-field
                v-model="ruleFirewallLocal.source_ip"
                label="Source IP"
                :error-messages="errors"
                required
              />
            </ValidationProvider>

            <ValidationProvider
              v-slot="{ errors }"
              name="Username"
              rules="required"
            >
              <v-text-field
                v-model="ruleFirewallLocal.username"
                label="Username"
                :error-messages="errors"
                required
              />
            </ValidationProvider>

            <ValidationProvider
              v-slot="{ errors }"
              name="Hostname"
              rules="required"
            >
              <v-text-field
                v-model="ruleFirewallLocal.hostname"
                label="Hostname"
                :error-messages="errors"
                required
              />
            </ValidationProvider>
          </v-card-text>

          <v-card-actions>
            <v-spacer />

            <v-btn
              text
              data-test="cancel-btn"
              @click="close"
            >
              Cancel
            </v-btn>

            <v-btn
              v-if="createRule"
              text
              data-test="create-btn"
              @click="passes(create)"
            >
              Create
            </v-btn>

            <v-btn
              v-else
              text
              data-test="edit-btn"
              @click="passes(edit)"
            >
              Edit
            </v-btn>
          </v-card-actions>
        </ValidationObserver>
      </v-card>
    </v-dialog>
  </fragment>
</template>

<script>

import {
  ValidationObserver,
  ValidationProvider,
} from 'vee-validate';

export default {
  name: 'FirewallRuleFormDialog',

  components: {
    ValidationProvider,
    ValidationObserver,
  },

  props: {
    firewallRule: {
      type: Object,
      required: false,
      default: Object,
    },
    createRule: {
      type: Boolean,
      required: true,
    },
  },

  data() {
    return {
      dialog: false,
      state: [{
        id: 'allow',
        name: 'allow',
      },
      {
        id: 'deny',
        name: 'deny',
      }],
      ruleFirewallLocal: [],
    };
  },

  computed: {
    isOwner() {
      return this.$store.getters['namespaces/owner'];
    },
  },

  async created() {
    await this.setLocalVariable();
  },

  async updated() {
    await this.setLocalVariable();
  },

  methods: {
    setLocalVariable() {
      if (this.createRule) {
        this.ruleFirewallLocal = {
          active: true,
          priority: '',
          action: '',
          source_ip: '',
          username: '',
          hostname: '',
        };
      } else {
        this.ruleFirewallLocal = { ...this.firewallRule };
      }
    },

    async create() {
      try {
        await this.$store.dispatch('firewallrules/post', this.ruleFirewallLocal);
        this.$store.dispatch('snackbar/showSnackbarSuccessAction', this.$success.firewallRuleCreating);
        this.update();
      } catch {
        this.$store.dispatch('snackbar/showSnackbarErrorAction', this.$errors.snackbar.firewallRuleCreating);
      }
    },

    async edit() {
      try {
        await this.$store.dispatch('firewallrules/put', this.ruleFirewallLocal);
        this.$store.dispatch('snackbar/showSnackbarSuccessAction', this.$success.firewallRuleEditing);
        this.update();
      } catch {
        this.$store.dispatch('snackbar/showSnackbarErrorAction', this.$errors.snackbar.firewallRuleEditing);
      }
    },

    update() {
      this.$emit('update');
      this.close();
    },

    close() {
      this.dialog = !this.dialog;
      this.$refs.obs.reset();
    },
  },
};
</script>
