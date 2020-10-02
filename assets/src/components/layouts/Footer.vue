<template>
    <footer class="footer">
        <div class="container">
            <div class="row footer-top">
                <div class="col-sm-5 col-lg-5">
                    <p class="padding-top-xsm">{{ description }}</p>
                </div>

                <div class="col-sm-6 col-lg-6 col-lg-offset-1">
                    <div class="row">
                        <div class="col-sm-12">
                            <h4>{{ other.title }}</h4>

                            <ul class="list-unstyled">
                                <li v-for="item in other.list">
                                    <a :href="item.link" v-title="item.title" target="_blank">
                                        <i :class="`fa fa-${item.icon}`"></i> {{ item.title }}
                                    </a>
                                </li>
                            </ul>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </footer>
</template>

<script>
    import title from '@/directives/title'
    export default {
        name: 'Footer',
        directives: {
            title
        },
        data() {
            return {
                description: 'article_create 是文章生成器',
                contactStyle: {
                    paddingRight: '8px'
                },

                other: {
                    title: '友情链接',
                    list: [
                        {
                            icon: 'thumbs-up',
                            title: '',
                            link: ''
                        },
                        {
                            icon: 'globe',
                            title: '',
                            link: ''
                        },
                        {
                            icon: 'globe',
                            title: '',
                            link: ''
                        },
                    ]
                }
            }
        }
    }
</script>

<style scoped>
    a:hover, a:focus { color: #e27575; transition: color .15s;}
</style>