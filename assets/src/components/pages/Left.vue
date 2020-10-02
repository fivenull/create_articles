<template>
    <div class="users-show">
        <div class="col-md-3 main-col">
            <div class="box">
                <div class="padding-md">
                    <div class="list-group text-center">
                        <router-link v-for="link in links" :key="link.name" :to="`${link.name}`" class="list-group-item">
                            <i :class="`text-md fa fa-${link.icon}`"></i>
                            {{ link.title }}
                        </router-link>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: 'Left',
        data() {
            return {
                links: [
                    {
                        name: '/histories',
                        title: '历史生成记录',
                        icon: 'list-alt'
                    },
                    {
                        name: '/categories',
                        title: '查看所有分类',
                        icon: 'list-alt'
                    }
                ]
            }
        }
    }
</script>

<style scoped>

</style>