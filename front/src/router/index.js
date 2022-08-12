import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/google-chart/GoogleChart.vue'
import AllTrafficView from '../views/trojanTraffic/AllTrafficView.vue'
import TrafficByGroup from '../views/trojanTraffic/TrafficByGroup.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/traffic',
    name: 'traffic',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: AllTrafficView,
    // children: [
    //   {
    //     path: '/traffic/group',
    //     name: 'trafficByGroup',
    //     component: TrafficByGroup
    //   }
    // ]
  },
  {
    path: '/traffic/group',
    name: 'trafficByGroup',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: TrafficByGroup
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
