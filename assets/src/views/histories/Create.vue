<template>
    <div class="blog-pages" style="padding-bottom:120px">
        <div class="col-md-9 panel">
            <div class="panel-body">
                <h2 class="text-center">生成文章</h2>
                <hr>
                <div data-validator-form>
                    <div class="form-group">
                        <button class="btn btn-success" type="button" @click="addArticleName">添加节点</button>
                        <button class="btn btn-primary mlr20" type="button" @click="getBody">生成预览</button>
                    </div>
                    <div class="form-group">
                        <label for='name'>标题</label>
                        <input v-model.trim="name" v-validator.required="{ name: '标题' }" type="text" class="form-control" placeholder="请填写标题">
                    </div>
                    <div class="form-inline pt pb" v-for="(category,index) in categoryList" :key="index">
                        <label for='c_id'>添加节点：</label>
                        <select name="c_id" v-validator.required="{ c_id: '分类' }" class="form-control select-article-ids">
                            <option value="">请选择</option>
                            <option v-for="items in category" :value="items.id">{{items.name}}</option>
                        </select>
                        <a class="admin mlr20" href="javascript:;" @click="removeArticleName(index)" title="删除"><i class="fa fa-trash-o"></i></a>
                    </div>
                    <div class="form-group pt">
                        <label for='body'>内容</label>
                        <textarea id="editor"></textarea>
<!--                        <div class="markdown-body" v-html="body"></div>-->
                    </div>
                    <br>
                    <div class="form-group">
                        <button class="btn btn-primary" type="submit" @click="saveBody">保存</button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import SimpleMDE from 'simplemde'
    import hljs from 'highlight.js'
    window.hljs = hljs
    export default {
        name: "Create",
        data() {
            return {
                name: '', // 节点标题
                body: '', // 节点内容
                order_by: 1,
                categoryList:[],//文章列表
                category:[],//文章
                msg: '', // 消息
                msgType: '', // 消息类型
                msgShow: false // 是否显示消息，默认不显示
            }
        },
        methods:{
            addArticleName(){
                this.categoryList.push(this.category)
            },
            removeArticleName(i){
                //判断长度，最短1
                console.log(this.categoryList.length)
                if(this.categoryList.length > 1) {
                    this.categoryList.splice(i ,1)
                }else{
                    alert('至少要保留一个哦')
                }
            },
            showMsg(msg, type = 'warning') {
                this.msg = msg
                this.msgType = type
                this.msgShow = true

                this.$nextTick(() => {
                    this.msgShow = true
                })
            },
            clearData() {
                this.name = ''
                this.order_by = 1
                this.c_id = ''
                this.simplemde.value('')
                this.simplemde.clearAutosavedValue()
            },
            //设置编辑器内容
            getBody(){
                let select = document.getElementsByName("c_id")
                let ids = [];


                if (!select[0].value){
                    alert('请选择节点')
                    return
                }
                for(let i=0; i< select.length; i++){
                    ids.push(select[i].value)
                }

                //提交输入的id ，获取详情
                //获取id，拼装ids
                let article_ids = ids.join(',');
                const article = {
                    article_ids
                }
                this.API.post("article/ids",'', article)
                .then(res =>{
                    this.simplemde.value(this.simplemde.value()+"\n"+res.data.data)
                })
            },
            //保存内容
            saveBody(){
                let body = this.simplemde.value()
                const name =this.name
                let select = document.getElementsByName("c_id")
                let ids = [];
                if (select.length <1){
                    alert('请选择节点')
                    return
                }
                for(let i=0; i< select.length; i++){
                    ids.push(select[i].value)
                }

                //提交输入的id ，获取详情
                //获取id，拼装ids
                let article_ids = ids.join(',');
                const article = {
                    article_ids,
                    name,
                    body
                }
                this.API.post("history/create",'', article)
                .then(res=>{
                    if(res.data.code == 0){
                        //生成成功，获取内容，进行跳转
                        alert("生成成功")
                        this.$router.push({ name: 'HistoryShow', params: { history_id:  res.data.data} })
                    }else{
                        alert(res.data.msg)
                    }
                })
            }
        },
        created(){
            this.API.get('article/name')
                .then(res => {
                    if(res.data.code == 0) {
                        let categorylists = res.data.data.items
                        if(categorylists.length < 1){
                            alert("请先创建分类")
                            this.$router.push("/category/create")
                        }
                        this.category = categorylists
                        this.categoryList.push(categorylists)
                    }
                })
        },
        mounted() {
            const simplemde = new SimpleMDE({
                element: document.querySelector('#editor'),
                placeholder: '请使用 Markdown 格式书写 ;-)，代码片段黏贴时请注意使用高亮语法。',
                spellChecker: false,
                autoDownloadFontAwesome: false,
                autosave: {
                    enabled: true,
                    uniqueId: 'mwxy_article'
                },
                renderingConfig: {
                    codeSyntaxHighlighting: true
                }
            })
            simplemde.codemirror.on('change', () => {
                this.body = simplemde.value()
            })

            this.simplemde = simplemde
            this.clearData()
        },
    }
</script>

<style scoped>
.mlr20{
    margin-left: 20px;
}
.pt{
    padding-top: 10px;
}
.pb{
    padding-top: 10px;
}
textarea { height: 200px; }
</style>