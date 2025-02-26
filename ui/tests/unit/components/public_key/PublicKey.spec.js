import Vuex from 'vuex';
import { mount, createLocalVue } from '@vue/test-utils';
import Vuetify from 'vuetify';
import publicKey from '@/components/public_key/PublicKey';

describe('Session', () => {
  const localVue = createLocalVue();
  const vuetify = new Vuetify();
  localVue.use(Vuex);

  let wrapper;

  const numberPublickeysEqualZero = 0;
  const numberPublickeysGreaterThanZero = 1;
  const isLoggedIn = true;

  const storeWithoutPublickeys = new Vuex.Store({
    namespaced: true,
    state: {
      numberPublickeys: numberPublickeysEqualZero,
      isLoggedIn,
    },
    getters: {
      'publickeys/getNumberPublicKeys': (state) => state.numberPublickeys,
      'auth/isLoggedIn': (state) => state.isLoggedIn,
    },
    actions: {
      'publickeys/refresh': () => {},
      'boxs/setStatus': () => {},
      'publickeys/resetPagePerpage': () => {},
      'snackbar/showSnackbarErrorLoading': () => {},
    },
  });

  const storeWithPublickeys = new Vuex.Store({
    namespaced: true,
    state: {
      numberPublickeys: numberPublickeysGreaterThanZero,
      isLoggedIn,
    },
    getters: {
      'publickeys/getNumberPublicKeys': (state) => state.numberPublickeys,
      'auth/isLoggedIn': (state) => state.isLoggedIn,
    },
    actions: {
      'publickeys/refresh': () => {},
      'boxs/setStatus': () => {},
      'publickeys/resetPagePerpage': () => {},
      'snackbar/showSnackbarErrorLoading': () => {},
    },
  });

  const storeWithoutPublickeysLogout = new Vuex.Store({
    namespaced: true,
    state: {
      numberPublickeys: numberPublickeysEqualZero,
      isLoggedIn: !isLoggedIn,
    },
    getters: {
      'publickeys/getNumberPublicKeys': (state) => state.numberPublickeys,
      'auth/isLoggedIn': (state) => state.isLoggedIn,
    },
    actions: {
      'publickeys/refresh': () => {},
      'boxs/setStatus': () => {},
      'publickeys/resetPagePerpage': () => {},
      'snackbar/showSnackbarErrorLoading': () => {},
    },
  });

  ///////
  // In this case, the rendering of the component that shows the
  // message when it does not have public key is tested.
  ///////

  describe('Without public key', () => {
    beforeEach(() => {
      wrapper = mount(publicKey, {
        store: storeWithoutPublickeys,
        localVue,
        stubs: ['fragment'],
        vuetify,
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

    ///////
    // Data and Props checking
    //////

    it('Compare data with the default and defined value', () => {
      expect(wrapper.vm.show).toEqual(true);
    });
    it('Process data in the computed', () => {
      expect(wrapper.vm.hasPublickey).toEqual(false);
      expect(wrapper.vm.showBoxMessage).toEqual(true);
    });

    //////
    // HTML validation
    //////

    it('Renders the template with components', () => {
      expect(wrapper.find('[data-test="publicKeyCreate-component"]').exists()).toBe(true);
      expect(wrapper.find('[data-test="boxMessagePublicKey-component"]').exists()).toBe(true);
    });
  });

  ///////
  // In this case, it is tested when there is already a registered
  // public key.
  ///////

  describe('Without public key', () => {
    beforeEach(() => {
      wrapper = mount(publicKey, {
        store: storeWithPublickeys,
        localVue,
        stubs: ['fragment'],
        vuetify,
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

    ///////
    // Data and Props checking
    //////

    it('Compare data with the default and defined value', () => {
      expect(wrapper.vm.show).toEqual(true);
    });
    it('Process data in the computed', () => {
      expect(wrapper.vm.hasPublickey).toEqual(true);
      expect(wrapper.vm.showBoxMessage).toEqual(false);
    });

    //////
    // HTML validation
    //////

    it('Renders the template with components', () => {
      expect(wrapper.find('[data-test="publicKeyCreate-component"]').exists()).toBe(true);
      expect(wrapper.find('[data-test="boxMessagePublicKey-component"]').exists()).toBe(false);
    });
  });

  ///////
  // In this case, purpose is to test the completion of the logout.
  // For this, the show variable must be false.
  ///////

  describe('Without public key', () => {
    beforeEach(() => {
      wrapper = mount(publicKey, {
        store: storeWithoutPublickeysLogout,
        localVue,
        stubs: ['fragment'],
        vuetify,
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

    ///////
    // Data and Props checking
    //////

    it('Compare data with the default and defined value', () => {
      expect(wrapper.vm.show).toEqual(false);
    });
    it('Process data in the computed', () => {
      expect(wrapper.vm.hasPublickey).toEqual(false);
      expect(wrapper.vm.showBoxMessage).toEqual(false);
    });

    //////
    // HTML validation
    //////

    it('Renders the template with components', () => {
      expect(wrapper.find('[data-test="publicKeyCreate-component"]').exists()).toBe(true);
      expect(wrapper.find('[data-test="boxMessagePublicKey-component"]').exists()).toBe(false);
    });
  });
});
