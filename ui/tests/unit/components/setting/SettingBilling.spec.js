import Vuex from 'vuex';
import { mount, createLocalVue } from '@vue/test-utils';
import SettingBilling from '@/components/setting/SettingBilling'
import Vuetify from 'vuetify';

describe('SettingBilling', () => {
  const localVue = createLocalVue();
  const vuetify = new Vuetify();
  localVue.use(Vuex);

  let wrapper;
  const owner = true;

  const billing = {
    active: true,
    current_period_end: 12121,
    customer_id: 'cus_123',
    subscription_id: 'subs_123',
    payment_method_id: 'pm_123',
  };

  const store = new Vuex.Store({
    namespaced: true,
    state: {
      stateBilling: billing.active,
      billing,
      owner,
    },
    getters: {
      'billing/active': (state) => state.stateBilling,
      'billing/get': (state) => state.billing,
      'namespaces/owner': (state) => state.owner,
    },
    actions: {
      'billing/getSubscription': () => {},
      'billing/cancelSubscription': () => {},
      'snackbar/showSnackbarSuccessAction': () => {},
      'snackbar/showSnackbarErrorAction': () => {},
      'snackbar/showSnackbarErrorDefault': () => {},
    },
  });

  beforeEach(() => {
    wrapper = mount(SettingBilling, {
      store,
      localVue,
      stubs: ['fragment'],
      vuetify,
      mocks: {
        $stripe: {
          elements: (args) => {
            return {
              create: (item) => {
                return {
                  mount: () => { return null; }
                };
              },
            };
          },
        },
      },
    });
  });

  ///////
  // Component Rendering
  //////

  it('Is a Vue instance', () => {
    expect(wrapper).toBeTruthy();
  });
  it('Renders the component', () => {
    expect(wrapper.html()).toMatchSnapshot();
  })

  ///////
  // Data and Props checking
  //////

  it('Compare data with default value', () => {
    expect(wrapper.vm.info).toEqual({});
    expect(wrapper.vm.cardItems).toEqual({});
    expect(wrapper.vm.renderIcon).toEqual(false);
  });
  it('Process data in the computed', () => {
    expect(wrapper.vm.active).toEqual(billing.active);
    expect(wrapper.vm.billing).toEqual(billing);
    expect(wrapper.vm.isOwner).toEqual(owner);
  });

  //////
  // HTML validation
  //////

  it('Renders the template with components', () => {
    expect(wrapper.find('[data-test="settingOwnerInfo-component"]').exists()).toBe(false);
    expect(wrapper.find('[data-test="subscriptionPaymentMethod-component"]').exists()).toBe(false);
    expect(wrapper.find('[data-test="updatePaymentMethod-component"]').exists()).toBe(true);
    expect(wrapper.find('[data-test="billingIcon-component"]').exists()).toBe(false);
  });
  it('Renders the template with data', async () => {
    expect(wrapper.find('[data-test="cancel-btn"]').exists()).toBe(true);
  });
})
