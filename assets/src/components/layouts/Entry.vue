<template>
    <div class="navbar-right">
        <ul class="nav navbar-nav github-login">
            <li>
                <a v-dropdown href="javascript:;">
                    <i class="fa fa-plus text-md"></i>
                </a>
                <ul class="dropdown-menu">
                    <li>
                        <router-link to="/category/create">
                            <i class="fa fa-paint-brush text-md"></i>
                            创建分类
                        </router-link>
                    </li>
                    <li>
                        <router-link to="/articles/create">
                            <i class="fa fa-paint-brush text-md"></i>
                            创建节点
                        </router-link>
                    </li>
                    <li>
                        <router-link to="/history/create">
                            <i class="fa fa-paint-brush text-md"></i>
                            创建文章
                        </router-link>
                    </li>
                </ul>
            </li>
        </ul>
    </div>
</template>

<script>
    export default {
        name: 'Entry'
    }
</script>

<style scoped>

</style>