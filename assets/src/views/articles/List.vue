<template>
    <div class="col-md-9 panel">
        <table class="table table-striped">
            <thead>
            <tr>
                <th scope="col">排序</th>
                <th scope="col">ID</th>
                <th scope="col">名称</th>
                <th scope="col">创建时间</th>
                <th scope="col">操作</th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="(article, index) in ArticleList" :key="index">
                <th scope="row">{{article.order_by}}</th>
                <th scope="row">{{article.id}}</th>
                <td>{{article.name}}</td>
                <td>{{article.create_date}}</td>
                <td>
                    <div class="actions">
                        <a @click="editArticle(category_id, article.id)" class="admin" href="javascript:;"><i class="fa fa-pencil-square-o"></i></a>
                        <a @click="deleteArticle(article.id)" class="admin" href="javascript:;"><i class="fa fa-trash-o"></i></a>
                    </div>
                </td>
            </tr>
            </tbody>
        </table>
    </div>
</template>

<script>
    export default {
        name : "Index",
        inject:['reload'],
        data(){
            return {
                ArticleList:[],
                isRouterAlive:true,
                category_id : this.$route.params.category_id,
            }
        },
        methods:{
            //获取列表
            _getArticleList(){
                this.API.get(`article/category/${this.category_id}`)
                    .then(res => {
                        console.log(res)
                        this.ArticleList = res.data.data.items
                    })
            },
            //删除
            deleteArticle(id){
                let _this = this
                //删除
                this.$swal({
                    text: '你确定要删除此内容吗?',
                    confirmButtonText: '删除'
                }).then((res) => {
                    if (res.value) {
                        _this.API.delete(`article/${id}`)
                            .then(res =>{
                                let data = res.data
                                if (data.code != 200){
                                    _this.reload()
                                }else{
                                    alert(data.msg)
                                }
                            })

                    }
                })
            },
            editArticle(category_id, article_id){
                this.$router.push({ name: 'ArticleEdit', params: { category_id: category_id, article_id:  article_id} })
            },
        },

        created(){
            this._getArticleList();
        }
    }
</script>

<style scoped>

</style>