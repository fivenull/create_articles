import Vue from 'vue'
import VueRouter from 'vue-router'
import routes from './routes'

Vue.use(VueRouter)


const router = new VueRouter({
  mode:'history',
  routes
})

router.beforeEach((to, from, next) => {
  const app = router.app
  const store = app.$options.store
  const categoryId = to.params.categoryId

  // if(
  //     (categoryId && store.getters.)
  // )else{
  //
  // }
  next()

})

export default router
