import SettingsService from '@/services/SettingsService.js'
export const state = () => ({
    settings: {},
 })
     
 export const mutations = {
   SET_SETTINGS(state, settings) {
     state.settings = settings;
   }
 }
 export const actions = {
  fetchSettings({ commit }) {
    return SettingsService.getSettings().then(response => {
      commit('SET_SETTINGS', response.data)
    })
  }
}