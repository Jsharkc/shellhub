import Vue from 'vue';
import getStats from '@/store/api/stats';

export default {
  namespaced: true,

  state: {
    stats: [],
  },

  getters: {
    stats: (state) => state.stats,
  },

  mutations: {
    setStats: (state, data) => {
      Vue.set(state, 'stats', data);
    },
  },

  actions: {
    get({
      commit,
    }) {
      return getStats()
        .then((response) => {
          commit('setStats', response.data);
        })
        .catch(() => Promise.reject());
    },
  },
};
