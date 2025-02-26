<template>
  <fragment>
    <v-row>
      <v-col
        v-if="!loggedInNamespace && isEnterprise && !isMobile"
      >
        <v-btn
          class="v-btn--active float-right mr-3"
          text
          small
          data-test="add-btn"
          @click="addNamespace"
        >
          Add Namespace
        </v-btn>
      </v-col>

      <v-col v-else>
        <v-menu
          v-show="displayMenu"
          :close-on-content-click="true"
          offset-y
          data-test="namespaceMenu-menu"
        >
          <template #activator="{ on }">
            <v-chip
              v-show="loggedInNamespace || isMobile"
              class="float-right"
              @click="openMenu"
              v-on="on"
            >
              <v-icon
                :size="defaultSize"
                left
              >
                mdi-server
              </v-icon>
              <div
                v-if="!isMobile"
                class="ml-1 mr-1"
                :size="defaultSize"
              >
                {{ namespace.name }}
              </div>
              <v-icon
                :size="defaultSize"
                right
              >
                mdi-chevron-down
              </v-icon>
            </v-chip>
          </template>

          <v-card>
            <div v-if="!isMobile">
              <v-subheader>Tenant ID</v-subheader>

              <v-list
                class="pt-0 pb-0 mx-2"
              >
                <v-list-item>
                  <v-row
                    justify="center"
                    align="center"
                  >
                    <v-chip>
                      <v-list-item-title>
                        <span data-test="tenantID-text">{{ tenant }}</span>
                      </v-list-item-title>
                      <v-icon
                        v-clipboard="tenant"
                        v-clipboard:success="() => {
                          $store.dispatch('snackbar/showSnackbarCopy', $copy.tenantId);
                        }"
                        right
                      >
                        mdi-content-copy
                      </v-icon>
                    </v-chip>
                  </v-row>
                </v-list-item>
              </v-list>
            </div>

            <v-divider />

            <v-list
              v-if="namespacesInList"
              class="pt-0"
            >
              <v-subheader>Namespaces</v-subheader>

              <v-list-item-group>
                <v-virtual-scroll
                  :height="adaptHeight"
                  item-height="50"
                  :items="availableNamespaces"
                >
                  <template #default="{ item }">
                    <v-list-item
                      :key="item.tenant_id"
                      @click="switchIn(item.tenant_id)"
                    >
                      <v-list-item-icon>
                        <v-icon>mdi-login</v-icon>
                      </v-list-item-icon>
                      <v-list-item-content>
                        <v-list-item-title>
                          {{ item.name }}
                        </v-list-item-title>
                      </v-list-item-content>
                    </v-list-item>
                  </template>
                </v-virtual-scroll>
              </v-list-item-group>
            </v-list>

            <v-divider />

            <v-list
              class="pt-0 pb-0"
            >
              <v-list-item
                v-show="isEnterprise"
                @click="dialog=!dialog"
              >
                <v-list-item-icon>
                  <v-icon>mdi-plus-box</v-icon>
                </v-list-item-icon>
                <v-list-item-content>
                  Create Namespace
                </v-list-item-content>
              </v-list-item>

              <v-divider />

              <v-list-item
                v-if="loggedInNamespace || !isMobile"
                to="/settings/namespace-manager"
              >
                <v-list-item-icon>
                  <v-icon>mdi-cog</v-icon>
                </v-list-item-icon>
                <v-list-item-content>
                  Settings
                </v-list-item-content>
              </v-list-item>
            </v-list>
          </v-card>
        </v-menu>
      </v-col>
      <NamespaceAdd
        :show.sync="dialog"
        :first-namespace="first"
        data-test="namespaceAdd-component"
      />
    </v-row>
  </fragment>
</template>

<script>
import NamespaceAdd from '@/components/app_bar/namespace/NamespaceAdd';

export default {
  name: 'NamespaceMenu',

  components: {
    NamespaceAdd,
  },

  props: {
    inANamespace: {
      type: Boolean,
      required: true,
    },
  },

  data() {
    return {
      model: true,
      dialog: false,
      displayMenu: false,
      defaultSize: 24,
      first: false,
    };
  },

  computed: {
    adaptHeight() {
      return this.availableNamespaces.length >= 3 ? 150 : this.availableNamespaces.length * 50;
    },

    isOwner() {
      return this.$store.getters['namespaces/owner'];
    },

    namespace() {
      return this.$store.getters['namespaces/get'];
    },

    namespaces() {
      return this.$store.getters['namespaces/list'];
    },

    namespacesInList() {
      return this.availableNamespaces.length > 0;
    },

    availableNamespaces() {
      return this.namespaces.filter((ns) => ns.tenant_id !== this.namespace.tenant_id);
    },

    loggedInNamespace() {
      return this.$props.inANamespace;
    },

    tenant() {
      return localStorage.getItem('tenant');
    },

    isEnterprise() {
      return this.$env.isEnterprise;
    },

    isMobile() {
      return this.$store.getters['mobile/isMobile'];
    },
  },

  watch: {
    dialog(value) {
      if (!value) {
        this.model = false;
      }
    },
  },

  async created() {
    await this.getNamespaces();
    if (this.inANamespace) {
      await this.getNamespace();
    }
  },

  methods: {
    addNamespace() {
      this.dialog = !this.dialog;
      this.first = true;
    },

    async openMenu() {
      if (!this.displayMenu) {
        await this.getNamespaces();
      }
      this.displayMenu = !this.displayMenu;
    },

    async getNamespace() {
      try {
        await this.$store.dispatch('namespaces/get', this.tenant);
        const isOwner = this.$store.getters['namespaces/get'].owner === this.$store.getters['auth/id'];
        this.$store.dispatch('namespaces/setOwnerStatus', isOwner);
      } catch (error) {
        switch (true) {
        case (error.response.status === 404): { // detects namespace inserted
          const namespaceFind = this.namespaces[0];
          if (this.tenant === '' && namespaceFind !== undefined) {
            this.switchIn(namespaceFind.tenant_id);
          }
          break;
        }
        case (error.response.status === 500 && this.tenant === null): {
          break;
        }
        default: {
          this.$store.dispatch('snackbar/showSnackbarErrorLoading', this.$errors.snackbar.namespaceLoad);
        }
        }
      }
    },

    async getNamespaces() {
      try {
        await this.$store.dispatch('namespaces/fetch');
      } catch (e) {
        switch (true) {
        case (!this.inANamespace && e.response.status === 403): { // dialog pops
          break;
        }
        case (e.response.status === 403): {
          this.$store.dispatch('snackbar/showSnackbarErrorAssociation');
          break;
        }
        default: {
          this.$store.dispatch('snackbar/showSnackbarErrorLoading', this.$errors.snackbar.namespaceList);
        }
        }
      }
    },

    async switchIn(tenant) {
      try {
        await this.$store.dispatch('namespaces/switchNamespace', {
          tenant_id: tenant,
        });

        const isOwner = this.$store.getters['namespaces/get'].owner === this.$store.getters['auth/id'];
        this.$store.dispatch('namespaces/setOwnerStatus', isOwner);

        window.location.reload();
      } catch {
        this.$store.dispatch('snackbar/showSnackbarErrorLoading', this.$errors.snackbar.namespaceSwitch);
      }
    },
  },
};
</script>
