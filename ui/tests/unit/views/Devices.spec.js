import Vuex from 'vuex';
import { shallowMount, createLocalVue } from '@vue/test-utils';
import Devices from '@/views/Devices';

describe('Terminal', () => {
  const localVue = createLocalVue();
  localVue.use(Vuex);

  let wrapper;

  beforeEach(() => {
    wrapper = shallowMount(Devices, {
      localVue,
      stubs: ['fragment'],
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

  //////
  // HTML validation
  //////

  it('Renders the template with components', () => {
    expect(wrapper.find('[data-test="deviceList-component"]').exists()).toBe(true);
  });
});
