<template>
  <fragment>
    <v-btn
      outlined
      data-test="show-btn"
      @click="show"
    >
      {{ typeOperation }}
    </v-btn>

    <!-- <div> -->
    <v-dialog
      v-model="dialog"
      max-width="600"
      @click:outside="dialog=!dialog"
    >
      <!-- <div> -->
      <v-card data-test="BillingDialogPaymentMethod-dialog">
        <v-card-title
          class="headline grey lighten-2 text-center"
          data-test="text-cardTitle"
        >
          {{ typeOperation | capitalizeFirstLetter }} payment method
        </v-card-title>

        <v-card-text>
          <v-card class="ma-4">
            <v-col>
              <div ref="card" />
            </v-col>
          </v-card>

          <div
            ref="card-element-errors"
            class="card-errors ma-4"
            role="alert"
          />

          <v-spacer />

          <v-row class="mt-2 mb-2">
            <v-spacer />
            <v-col
              md="auto"
              class="ml-auto"
            />
          </v-row>
        </v-card-text>

        <v-card-actions
          class="mr-5 pb-5"
        >
          <v-spacer />

          <v-btn
            outlined
            data-test="confirm-btn"
            @click="doAction()"
          >
            confirm
          </v-btn>
        </v-card-actions>
      </v-card>
      <!-- </div> -->
    </v-dialog>
    <!-- </div> -->
  </fragment>
</template>

<script>

import capitalizeFirstLetter from '@/components/filter/word';

export default {
  name: 'BillingDialogPaymentMethod',

  filters: { capitalizeFirstLetter },

  props: {
    typeOperation: {
      type: String,
      default: 'subscription',
      validator: (value) => ['subscription', 'update'].includes(value),
    },
  },

  data() {
    return {
      dialog: false,
      card: null,
      elementError: null,
      elements: null,
    };
  },

  methods: {
    show() {
      this.dialog = !this.dialog;
      this.$nextTick(async () => {
        await this.mountStripeElements();
      });
    },

    displayError(e) {
      if (e.error) {
        this.elementError.textContent = e.error.message;
      } else {
        this.elementError.textContent = '';
      }
    },

    doAction() {
      switch (this.typeOperation) {
      case 'subscription':
        this.subscriptionPaymentMethod();
        break;
      case 'update':
        this.updatePaymentMethod();
        break;
      default:
      }
    },

    async mountStripeElements() {
      this.elements = this.$stripe.elements();
      this.card = this.elements.create('card');
      this.card.mount(this.$refs.card);
      this.elementError = this.$refs['card-element-errors'];
      this.card.on('change', (ev) => {
        this.displayError(ev);
      });
    },

    async subscriptionPaymentMethod() {
      const result = await this.$stripe.createPaymentMethod({
        type: 'card',
        card: this.card,
      });

      if (result.error) {
        this.displayError(result.error);
      } else {
        try {
          await this.$store.dispatch('billing/subscritionPaymentMethod', {
            payment_method_id: result.paymentMethod.id,
          });
          this.$store.dispatch('snackbar/showSnackbarSuccessAction', this.$success.subscription);
        } catch {
          this.$store.dispatch('snackbar/showSnackbarErrorAction', this.$errors.snackbar.subscription);
        }
      }
    },

    async updatePaymentMethod() {
      const result = await this.$stripe.createPaymentMethod({
        type: 'card',
        card: this.card,
      });

      if (result.error) {
        this.displayError(result.error);
      } else {
        try {
          await this.$store.dispatch('billing/updatePaymentMethod', {
            payment_id: result.paymentMethod.id,
          });
          this.$store.dispatch('snackbar/showSnackbarSuccessAction', this.$success.updateSubscription);
        } catch {
          this.$store.dispatch('snackbar/showSnackbarErrorAction', this.$errors.snackbar.updateSubscription);
        }
      }
    },
  },
};

</script>
