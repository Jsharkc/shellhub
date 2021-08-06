import Vuex from 'vuex';
import { shallowMount, createLocalVue } from '@vue/test-utils';
import BillingDialogPaymentMethod from '@/components/billing/BillingDialogPaymentMethod';

describe('BillingDialogPaymentMethod', () => {
  const localVue = createLocalVue();
  localVue.use(Vuex);

  let wrapper;

  let typeOperation = 'subscription';

  const store = new Vuex.Store({
    namespaced: true,
    state: {
    },
    getters: {
    },
    actions: {
      'billing/subscritionPaymentMethod': () => {},
      'billing/updatePaymentMethod': () => {},
      'snackbar/showSnackbarSuccessAction': () => {},
      'snackbar/showSnackbarErrorAction': () => {},
    },
  });

  ///////
  // In this case, it's testing the button rendering.
  ///////

  describe('Button', () => {
    beforeEach(() => {
      wrapper = shallowMount(BillingDialogPaymentMethod, {
        localVue,
        stubs: ['fragment'],
        propsData: { typeOperation },
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
    });
  
    /////
    // Data and Props checking
    ////
  
    it('Receive data in props', () => {
      expect(wrapper.vm.typeOperation).toBe(typeOperation);
    });
    it('Compare data with default value', () => {
      expect(wrapper.vm.dialog).toEqual(false);
      expect(wrapper.vm.card).toEqual(null);
      expect(wrapper.vm.elementError).toEqual(null);
      expect(wrapper.vm.elements).toEqual(null);
    });
  
    //////
    // HTML validation
    //////
  
    it('Renders the template with data', async () => {
      expect(wrapper.find('[data-test="show-btn"]').exists()).toBe(true);
      // expect(wrapper.find('[data-test="BillingDialogPaymentMethod-dialog"]').exists()).toBe(false);
      // expect(wrapper.find('[data-test="confirm-btn"]').exists()).toBe(false);
    });
  });

  ///////
  // In this case, it's testing the subscription.
  ///////

  describe('Dialog Subscription', () => {
    beforeEach(() => {
      wrapper = shallowMount(BillingDialogPaymentMethod, {
        localVue,
        stubs: ['fragment'],
        propsData: { typeOperation },
      });

      wrapper.setData({ dialog: true });
    });
  
    ///////
    // Component Rendering
    //////
  
    it('Is a Vue instance', () => {
      expect(wrapper).toBeTruthy();
    });
    it('Renders the component', () => {
      expect(wrapper.html()).toMatchSnapshot();
    });
  
    /////
    // Data and Props checking
    ////
  
    it('Receive data in props', () => {
      expect(wrapper.vm.typeOperation).toBe(typeOperation);
    });
    it('Compare data with default value', () => {
      expect(wrapper.vm.dialog).toEqual(true);
      expect(wrapper.vm.card).toEqual(null);
      expect(wrapper.vm.elementError).toEqual(null);
      expect(wrapper.vm.elements).toEqual(null);
    });
  
    //////
    // HTML validation
    //////
  
    it('Renders the template with data', async () => {
      // expect(wrapper.find('[data-test="show-btn"]').exists()).toBe(false);
      expect(wrapper.find('[data-test="BillingDialogPaymentMethod-dialog"]').exists()).toBe(true);
      expect(wrapper.find('[data-test="confirm-btn"]').exists()).toBe(true);
      // expect(wrapper.find('[data-test="text-cardTitle"]').text()).toEqual('Subscription payment method');
    });
  });

  ///////
  // In this case, it's testing the update subscription.
  ///////

  describe('Dialog Update', () => {
    typeOperation = 'update';

    beforeEach(() => {
      wrapper = shallowMount(BillingDialogPaymentMethod, {
        localVue,
        stubs: ['fragment'],
        propsData: { typeOperation },
      });

      wrapper.setData({ dialog: true });
    });
  
    ///////
    // Component Rendering
    //////
  
    it('Is a Vue instance', () => {
      expect(wrapper).toBeTruthy();
    });
    it('Renders the component', () => {
      expect(wrapper.html()).toMatchSnapshot();
    });
  
    /////
    // Data and Props checking
    ////
  
    it('Receive data in props', () => {
      expect(wrapper.vm.typeOperation).toBe(typeOperation);
    });
    it('Compare data with default value', () => {
      expect(wrapper.vm.dialog).toEqual(true);
      expect(wrapper.vm.card).toEqual(null);
      expect(wrapper.vm.elementError).toEqual(null);
      expect(wrapper.vm.elements).toEqual(null);
    });
  
    //////
    // HTML validation
    //////
  
    it('Renders the template with data', async () => {
      // expect(wrapper.find('[data-test="show-btn"]').exists()).toBe(false);
      expect(wrapper.find('[data-test="BillingDialogPaymentMethod-dialog"]').exists()).toBe(true);
      expect(wrapper.find('[data-test="confirm-btn"]').exists()).toBe(true);
      expect(wrapper.find('[data-test="text-cardTitle"]').text()).toEqual('Update payment method');
    });
  });
});
