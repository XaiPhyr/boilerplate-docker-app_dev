import api from './__init__';

export default {
  async get() {
    return api('users');
  },
};
