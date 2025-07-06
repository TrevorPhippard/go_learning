import { mount } from '@vue/test-utils';
import Login from '../Login.vue';
import { describe, it, expect } from 'vitest';

describe('Login.vue', () => {
  it('renders login form', () => {
    const wrapper = mount(Login);
    expect(wrapper.find('form').exists()).toBe(true);
  });
});
