<template>
    <div class="blog-pages"  style="overflow-y:auto; overflow-x:auto; max-width:1080px; height:900px;">
        <Message :show.sync="msgShow" :type="msgType" :msg="msg"/>
        <div class="col-md-9 panel">
            <div class="panel-body">
                <h2 class="text-center">编辑分类</h2>
                <hr>
                <div data-validator-form>
                    <div class="form-group">
                        <label for="name">标题:</label>
                        <input v-model.trim="name" v-validator.required="{ name: '标题' }" type="text" class="form-control" placeholder="请填写标题">
                    </div>
                    <div class="form-group">
                        <label for="info">简介:</label>
                        <input v-model.trim="info" v-validator.required="{ info: '标题' }" type="text" class="form-control" placeholder="请填写简介">
                    </div>
                    <div class="form-group">
                        <label for="order_by">排序:</label>
                        <input v-model.trim="order_by" v-validator.required="{ info: '排序' },{ regex:/^[1,255]$/, error:'只能为1-255的整数' }" type="number" class="form-control" placeholder="请填写排序">
                    </div>
                    <br>
                    <div class="form-group">
                        <button class="btn btn-primary" type="submit" @click="categoryCreate">发 布</button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name :"Create",
        data (){
          return {
              name :"",
              info:"",
              order_by: 1,
              msg: '', // 消息
              msgType: '', // 消息类型
              msgShow: false // 是否显示消息，默认不显示
          }
        },
        methods: {
            categoryCreate(){
              const name = this.name
              const info = this.info
              const order_by = Number(this.order_by)
                console.log(order_by)
                if (name !== '' && info !== '' && order_by > 0) {
                    const category = {
                        name,
                        info,
                        order_by,
                    }
                    this.API.post("category/create", '', category)
                    .then(res =>{
                        if(res.data.code == 0){
                            this.$router.push('/categories')
                        }else{
                            this.showMsg(res.data.msg, 'danger')
                        }
                    })
                }
            },
            showMsg(msg, type = 'warning') {
                this.msg = msg
                this.msgType = type
                this.msgShow = false

                this.$nextTick(() => {
                    this.msgShow = true
                })
            }
        }
    }
</script>

<style scoped>
</style>