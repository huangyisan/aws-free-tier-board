import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/google-chart/GoogleChart.vue'
import AllTraffic from '../views/trojanTraffic/AllTrafficView.vue'
import TrafficByGroup from '../views/trojanTraffic/AllTrafficByGroup.vue'
import FilterTraffic from '../views/trojanTraffic/FilterTraffic.vue'

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
    component: AllTraffic,
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
  },
  {
    path: '/traffic/filter',
    name: 'trafficFilter',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: FilterTraffic
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
