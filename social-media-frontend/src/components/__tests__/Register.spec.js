import { mount } from '@vue/test-utils';
import Register from '../Register.vue';
import { describe, it, expect } from 'vitest';

describe('Register.vue', () => {
  it('renders the form inputs', () => {
    const wrapper = mount(Register);
    expect(wrapper.find('input[type="text"]').exists()).toBe(true);
    expect(wrapper.find('input[type="email"]').exists()).toBe(true);
    expect(wrapper.findAll('input[type="password"]').length).toBe(2);
  });

  it('shows error when passwords do not match', async () => {
    const wrapper = mount(Register);
    await wrapper.find('input[type="text"]').setValue('testuser');
    await wrapper
      .find('input[type="email"]')
      .setValue('test@example.com');
    await wrapper
      .findAll('input[type="password"]')[0]
      .setValue('123');
    await wrapper
      .findAll('input[type="password"]')[1]
      .setValue('456');
    await wrapper.find('form').trigger('submit.prevent');

    expect(wrapper.html()).toContain('Passwords do not match');
  });
});
