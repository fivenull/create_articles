<template>
    <div class="navbar navbar-default topnav">
        <div class="container">
            <div class="navbar-header">
                <button type="button" class="navbar-toggle" @click="toggleNav">
                    <span class="sr-only">Toggle navigation</span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                </button>


                <router-link to="/" class="navbar-brand">
                    <span class="title">{{ logo.title }}</span>
                    <img :src="logo.src" :alt="logo.title">
                </router-link>

            </div>

            <div id="top-navbar-collapse" :class="['collapse', 'navbar-collapse', { in: showCollapsedNav }]">
                <ul class="nav navbar-nav">
                    <li v-for="(item, index) in navList" :class="{ active: index === activeNavIndex }">
                        <router-link to="/" @click="changeNavIndex(index)">{{ item }}</router-link>
                    </li>
                </ul>
                <!-- 入口组件 -->
                <div class="navbar-right">
                    <Entry/>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import Entry from "@/components/layouts/Entry"
    export default {
        name: 'Header',
        components:{
            Entry
        },
        data() {
            return {
                logo: {
                    src: `${this.uploadsUrl}pc/images/logo_new.png`,
                    title: 'Learnku Vue.js'
                },
                navList: ['节点生成'],
                activeNavIndex: 0,
                showCollapsedNav: false
            }
        },
        beforeCreate() {
            this.uploadsUrl = 'http://www.gzmiwi.com/'
        },
        methods: {
            changeNavIndex(index) {
                this.activeNavIndex = index
            },
            toggleNav() {
                this.showCollapsedNav = !this.showCollapsedNav
            }
        }
    }
</script>

<style scoped>
    .title { display: none;}
    .navbar {
        border-radius: 0px !important;
        }
    .navbar-default .navbar-nav > .active > a { background: rgba(0,0,0,.03);}
</style>