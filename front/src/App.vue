<template>
  <v-app id="inspire">
    <v-navigation-drawer v-model="drawer" app>
      <!--  -->
      <v-list rounded>
        <v-subheader>Cost Report</v-subheader>
        <v-list-item-group>
          <!-- <v-list-item>
            <v-list-item-icon>
              <v-icon>mdi-home</v-icon>
            </v-list-item-icon>

            <v-list-item-title>Home</v-list-item-title>
          </v-list-item> -->

          <v-list-group :value="true" prepend-icon="mdi-wan">
            <template v-slot:activator>
              <v-list-item-title>Trojan</v-list-item-title>
            </template>

            <v-list-group :value="true" no-action sub-group>
              <template v-slot:activator>
                <v-list-item-content>
                  <v-list-item-title>DashBoard</v-list-item-title>
                </v-list-item-content>
              </template>
              <v-col>
                <v-list-item
                  v-for="(item, i) in trojanTraffic"
                  :key="i"
                  link
                  exact
                  :to="item.to"
                >
                  <v-list-item-icon>
                    <v-icon v-text="item.icon"></v-icon>
                  </v-list-item-icon>
                  <v-list-item-title v-text="item.title"></v-list-item-title>
                </v-list-item>
              </v-col>
            </v-list-group>

            <!-- <v-list-group no-action sub-group>
              <template v-slot:activator>
                <v-list-item-content>
                  <v-list-item-title>Users</v-list-item-title>
                </v-list-item-content>
              </template>

              <v-list-item v-for="([title, icon], i) in cruds" :key="i" link>
                <v-list-item-title v-text="title"></v-list-item-title>

                <v-list-item-icon>
                  <v-icon v-text="icon"></v-icon>
                </v-list-item-icon>
              </v-list-item>
            </v-list-group> -->
          </v-list-group>
        </v-list-item-group>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar app>
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>

      <v-toolbar-title>Application</v-toolbar-title>
    </v-app-bar>

    <v-main>
      <!--  -->
      <router-view></router-view>
    </v-main>
  </v-app>
</template>


<script>
export default {
  name: "App",

  data: () => ({
    trojanTraffic: [
      { title: "All Traffic", icon: "mdi-download", to: "/traffic" },
      { title: "Traffic By Group", icon: "mdi-download", to: "/traffic/group" },
      // ["All Traffic", "mdi-download"],
      // ["Settings", "mdi-cog-outline"],
    ],
    cruds: [
      ["Create", "mdi-plus-outline"],
      ["Read", "mdi-file-outline"],
      ["Update", "mdi-update"],
      ["Delete", "mdi-delete"],
    ],
    drawer: null,
    selectedItem: 1,
    // items: [
    //   { text: "home", icon: "mdi-home", to: "/" },
    //   { text: "Trojan Traffic", icon: "mdi-wan", to: "trojanTraffic" },
    // ],
  }),
};
</script>
