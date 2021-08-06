<template>
  <v-container>
    <v-row
      align="center"
      justify="center"
      class="mt-4"
    >
      <v-col
        sm="8"
      >
        <SettingOwnerInfo
          :is-owner="isOwner"
          data-test="settingOwnerInfo-component"
        />

        <div>
          <v-row>
            <v-col>
              <h3>
                Subscription info
              </h3>
            </v-col>

            <v-spacer />

            <v-col
              v-if="isOwner && !active"
              md="auto"
              class="ml-auto"
            >
              <PaymentMethod
                type-operation="subscription"
                data-test="subscriptionPaymentMethod-component"
              />
            </v-col>
          </v-row>

          <div class="mt-6 pl-4 pr-4">
            <div v-if="isOwner && !active">
              <p>
                Plan: <b> Free </b>
              </p>

              <p>
                Description: you can only add 3 devices..
              </p>
            </div>

            <div v-else-if="isOwner && active">
              <p>
                Plan: <b> {{ info.description }} </b>
              </p>

              <p>
                Description: In this plan, the amount is charged according to the number of
                devices used.
              </p>
            </div>
          </div>
        </div>

        <div
          v-if="isOwner && active"
          class="mt-4 mb-4"
        >
          <v-divider />
          <v-divider />

          <div class="mt-6">
            <v-row>
              <v-col>
                <h3>
                  Next bill
                </h3>
              </v-col>
            </v-row>

            <div class="mt-6 pl-4 pr-4">
              <p>
                Date: <b> {{ info.periodEnd | formatDateWithoutDayAndHours }} </b>
              </p>

              <p>
                Estimated total: <b> $ {{ info.nextPaymentDue }} </b>
              </p>
            </div>
          </div>

          <v-divider />
          <v-divider />

          <div class="mt-6">
            <v-row>
              <v-col>
                <h3>
                  Payment method details
                </h3>
              </v-col>

              <v-spacer />

              <v-col
                md="auto"
                class="ml-auto"
              >
                <PaymentMethod
                  type-operation="update"
                  data-test="updatePaymentMethod-component"
                />
              </v-col>
            </v-row>

            <div class="mt-5 pl-4 pr-4">
              <p>
                <BillingIcon
                  v-if="renderIcon"
                  :icon-name="cardItems.brand"
                  data-test="billingIcon-component"
                />
                {{ cardItems.expMonth }}/{{ cardItems.expYear }} - {{ cardItems.last4 }}
              </p>
            </div>
          </div>

          <v-divider />
          <v-divider />

          <div class="mt-6">
            <v-row>
              <v-col>
                <h3>
                  Cancel Subscription
                </h3>
              </v-col>
            </v-row>

            <div class="mt-2 pl-4">
              <v-row>
                <v-col>
                  <p>
                    When canceling subscription, you may lose access to devices.
                  </p>
                </v-col>

                <v-spacer />

                <v-col
                  md="auto"
                  class="ml-auto"
                >
                  <v-btn
                    color="red darken-1"
                    outlined
                    data-test="cancel-btn"
                    @click="cancelSubscription()"
                  >
                    Cancel
                  </v-btn>
                </v-col>
              </v-row>
            </div>
          </div>
        </div>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>

import SettingOwnerInfo from '@/components/setting/SettingOwnerInfo';
import PaymentMethod from '@/components/billing/BillingDialogPaymentMethod';
import BillingIcon from '@/components/billing/BillingIcon';

import { formatDateWithoutDayAndHours } from '@/components/filter/date';

export default {
  name: 'SettingBilling',

  components: {
    SettingOwnerInfo,
    PaymentMethod,
    BillingIcon,
  },

  filters: { formatDateWithoutDayAndHours },

  data() {
    return {
      card: null,
      elements: null,
      info: {},
      cardItems: {},
      renderIcon: false,
    };
  },

  computed: {
    active() {
      return this.$store.getters['billing/active'];
    },

    billing() {
      return this.$store.getters['billing/get'];
    },

    isOwner() {
      return this.$store.getters['namespaces/owner'];
    },
  },

  created() {
    if (this.isOwner && !this.active) {
      this.mountStripeElements();
    }
    this.getSubscriptionInfo();
  },

  methods: {
    mountStripeElements() {
      this.elements = this.$stripe.elements();
      this.card = this.elements.create('card');
      this.card.mount(this.$refs.card);
    },

    getInfo(data) {
      const li = data.latest_invoice;
      const ui = data.upcoming_invoice;
      const description = data.product_description;
      const { card } = data;

      return {
        info: {
          periodEnd: this.billing.current_period_end,
          description,
          latestPaymentDue: li.amount_due,
          latestPaymentPaid: li.amount_paid,
          nextPaymentDue: ui.amount_due,
          nextPaymentPaid: ui.amount_paid,
        },
        card: {
          brand: card.brand,
          expYear: card.exp_year,
          expMonth: card.exp_month,
          last4: card.last4,
        },
      };
    },

    async getSubscriptionInfo() {
      if (this.active) {
        try {
          const data = await this.$store.dispatch('billing/getSubscription');
          const { info, card } = this.getInfo(data);
          this.info = info;
          this.cardItems = card;

          this.renderIcon = true;
        } catch {
          // this.renderIcon = false;
          this.$store.dispatch('snackbar/showSnackbarErrorDefault');
        }
      }
    },

    async cancelSubscription() {
      try {
        await this.$store.dispatch('billing/cancelSubscription');
        this.$store.dispatch('snackbar/showSnackbarSuccessAction', this.$success.cancelSubscription);
      } catch {
        this.$store.dispatch('snackbar/showSnackbarErrorAction', this.$errors.snackbar.cancelSubscription);
      }
    },
  },
};

</script>
