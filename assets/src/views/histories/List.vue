<template>
    <div class="col-md-9 panel">
        <table class="table table-striped">
            <thead>
            <tr>
                <th scope="col">ID</th>
                <th scope="col">名称</th>
                <th scope="col">创建时间</th>
                <th scope="col">操作</th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="(history, index) in historyList" :key="index">
                <th scope="row">{{history.id}}</th>
                <td>{{history.name}}</td>
                <td>{{history.create_date}}</td>
                <td>
                    <div class="actions">
                        <a  class="admin" title="查看" href="javascript:;" @click="showHistory(history.id)"><i class="fa fa-search"></i></a>
                        <span class="count_seperator">|</span>
                        <a  class="admin" title="删除" href="javascript:;" @click="deleteHistory(history.id)"><i class="fa fa-trash-o"></i></a>
                    </div>
                </td>
            </tr>
            </tbody>
        </table>
        <div class="panel-footer text-right remove-padding-horizontal pager-footer">
            <Pagination :currentPage="currentPage" :total="total" :pageSize="pageSize" :onPageChange="changePage" />
        </div>
    </div>
</template>

<script>
    import { mapState } from 'vuex'
    export default {
        name : "Lists",
        inject:['reload'],
        data(){
            return {
                historyList:[],
                isRouterAlive:true,
                page:1,
                order_by:'created_at',
                order:'DESC',
                total: 0, // 文章总数
                pageSize: 10, // 每页条数
            }
        },
        computed: {
            ...mapState([
                'auth',
                'user'
            ]),
            // 当前页，从查询参数 page 返回
            currentPage() {
                return parseInt(this.$route.query.page) || 1
            }
        },
        methods:{
            //获取列表
            _getHistoryList(page = 1){
                const history = {
                    page:page,
                    order_by:this.order_by,
                    order:this.order,
                }
                this.API.get('history', history)
                    .then(res => {
                        console.log(res)
                        this.total = res.data.data.total
                        this.historyList = res.data.data.items
                    })
            },
            //删除
            deleteHistory(id){
                let _this = this
                //删除
                this.$swal({
                    text: '你确定要删除此内容吗?',
                    confirmButtonText: '删除'
                }).then((res) => {
                    if (res.value) {
                        _this.API.delete(`history/${id}`)
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
           showHistory(history_id){
                this.$router.push({ name: 'HistoryShow', params: { history_id:  history_id} })
            },
            // 回调，组件的当前页改变时调用
            changePage(page) {
                console.log(page);
                // 在查询参数中混入 page，并跳转到该地址
                // 混入部分等价于 Object.assign({}, this.$route.query, { page: page })
                this._getHistoryList(page)
                this.$router.push({ query: { ...this.$route.query, page } })
            }
        },

        created(){
            this._getHistoryList();
        }
    }
</script>

<style scoped>

</style>