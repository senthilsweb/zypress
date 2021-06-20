import DocsService from '@/services/DocsService.js'
export const state = () => ({
    docs: [],
 })
     
 export const mutations = {
   SET_DOCS(state, docs) {
     state.docs = docs;
   }
 }
 export const actions = {
  fetchDocs({ commit }) {
    return DocsService.getDocs().then(response => {
      commit('SET_DOCS', response.data)
    })
  }
}