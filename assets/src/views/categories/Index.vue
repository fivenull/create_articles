<template>
    <div class="col-md-9 panel">
        <table class="table table-striped">
            <thead>
            <tr>
                <th scope="col">ID</th>
                <th scope="col">名称</th>
                <th scope="col">简介</th>
                <th scope="col">创建时间</th>
                <th scope="col">操作</th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="(category, index) in categoryList" :key="index">
                <th scope="row">{{category.id}}</th>
                <td>{{category.name}}</td>
                <td>{{category.info}}</td>
                <td>{{category.create_date}}</td>
                <td>
                    <div class="actions">
                        <a @click="ArticleList(category.id)" class="admin" href="javascript:;" title="查看下属节点"><i class="fa fa-folder"></i></a>
                        <span class="count_seperator">|</span>
                        <a @click="editCategory(category.id)" class="admin" href="javascript:;" title="编辑分类"><i class="fa fa-pencil-square-o"></i></a>
                        <span class="count_seperator">|</span>
                        <a @click="deleteCategory(category.id)" class="admin" href="javascript:;" title="删除分类"><i class="fa fa-trash-o"></i></a>
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
            categoryList:[],
            isRouterAlive:true,
        }
    },
    methods:{
        //获取列表
        _getCategoryList(){
            console.log(this)
            this.API.get('category/lists')
            .then(res => {
                console.log(res);
                this.categoryList = res.data.data.items
            })
        },
        //删除
        deleteCategory(id){
            let _this = this
            //删除
            this.$swal({
                text: '你确定要删除此内容吗?',
                confirmButtonText: '删除'
            }).then((res) => {
                if (res.value) {
                    // id = 100
                    _this.API.delete(`category/${id}`)
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
        editCategory(id){
            this.$router.push({ name: 'CategoryEdit', params: { category_id: id } })
        },
        //查看下属节点内容
        ArticleList(id){
            this.$router.push({ name: 'ArticleList', params: { category_id: id } })
        }

    },

    created(){
        this._getCategoryList();
    }
}
</script>

<style scoped>

</style>