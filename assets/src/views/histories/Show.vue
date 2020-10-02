<template>
    <div class="blog-pages" style="padding-bottom:120px">
        <div class="col-md-9 panel">
            <div class="panel-body">
                <h2 class="text-center">{{name}}</h2>
                <hr>
                <div data-validator-form>
                    <div class="form-group pt">
                        <div class="markdown-body" v-html="body">内容</div>
                    </div>
                    <br>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import SimpleMDE from 'simplemde'
    import hljs from 'highlight.js'

    export default {
        name: "HistoryShow",
        data(){
            return {
                name : '',
                body: '',
            }
        },
        created(){
            let history_id = this.$route.params.history_id
            console.log(history_id)
            this.API.get(`history/get/${history_id}`)
            .then(res=>{
                console.log(res)
                if(res.data.code == 0){
                    this.name = res.data.data.name
                    let body = res.data.data.body
                    this.body = SimpleMDE.prototype.markdown(body)

                    this.$nextTick(() => {
                        this.$el.querySelectorAll('pre code').forEach((el) => {
                            hljs.highlightBlock(el)
                        })
                    })
                }else{
                    alert(res.data.msg)
                }
            })
        }

    }
</script>

<style scoped>

</style>