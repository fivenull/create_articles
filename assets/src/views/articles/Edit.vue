<template>
    <div class="blog-pages"  style="overflow-y:auto; overflow-x:auto; max-width:1080px; height:1000px;">
        <div class="col-md-9 panel">
            <div class="panel-body">
                <h2 class="text-center">编辑节点</h2>
                <hr>
                <div data-validator-form>
                    <div class="form-group">
                        <label for='name'>标题</label>
                        <input v-model.trim="name" v-validator.required="{ name: '标题' }" type="text" class="form-control" placeholder="请填写标题">
                    </div>
                    <div class="form-group">
                        <label for='c_id'>所属分类</label>
                        <select v-model.Number="c_id" v-validator.required="{ c_id: '分类' }" class="form-control">
                            <option value="">请选择</option>
                            <option v-for="category in categoryList" :value="category.id">{{category.name}}</option>
                        </select>
                    </div>
                    <div class="form-group">
                        <label for='order_by'>排序</label>
                        <input v-model.trim="order_by" v-validator.required="{ order_by: '标题' },{ regex:/[1-255]/, error:'只能为1-255的整数' }" type="number" class="form-control" placeholder="请填写排序" >
                    </div>
                    <div class="form-group">
                        <label for='body'>内容</label>
                        <textarea id="editor"></textarea>
<!--                        <div class="markdown-body" v-html="body"></div>-->
                    </div>
                    <br>
                    <div class="form-group">
                        <button class="btn btn-primary" type="submit" @click="put">发 布</button>
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
        name: 'Edit',
        data() {
            return {
                article_id: 0, //节点id
                name: '', // 节点标题
                body: '', // 节点内容
                c_id: 0, // 所属分类id
                order_by: 1,
                categoryList:[],//分类列表
                msg: '', // 消息
                msgType: '', // 消息类型
                msgShow: false // 是否显示消息，默认不显示
            }
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
            // this.clearData()
        },
        created(){
            this.API.get('category/lists')
                .then(res => {
                    if(res.data.code == 0) {
                        let categorylists = res.data.data.items
                        if(categorylists.length < 1){
                            alert("请先创建分类")
                            this.$router.push("/category/create")
                        }
                        this.categoryList = categorylists
                        this.API.get(`article/get/${this.$route.params.article_id}`)
                            .then(res => {
                                if (res.data.code == 0){
                                    this.c_id = Number(this.$route.params.category_id)
                                    this.name = res.data.data.name
                                    this.simplemde.value(res.data.data.body)
                                    this.order_by = res.data.data.order_by
                                }
                            })
                    }
                })

        },
        methods: {
            put() {
                const name = this.name
                const body = this.body
                const c_id = this.c_id
                const order_by = this.order_by
                if (name !== '' && body.trim() !== '' && c_id > 0) {
                    const article = {
                        name,
                        body,
                        c_id,
                        order_by,
                    }
                    console.log(article)
                    //提交
                    this.API.put(`article/${this.$route.params.article_id}`, '', article)
                        .then(res => {
                            if(res.data.code == 0){
                                alert('更新成功')
                                this.$router.push({ name: 'ArticleList', params: { category_id: c_id } })
                            }else{
                                this.showMsg(res.data.msg)
                            }
                        })
                    // this.clearData()
                }
            },
            clearData() {
                this.title = ''
                //ls.removeItem('smde_title')
                this.simplemde.value('')
                this.simplemde.clearAutosavedValue()
            },
            showMsg(msg, type = 'warning') {
                this.msg = msg
                this.msgType = type
                this.msgShow = false

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
        }
    }
</script>

<style scoped>
    .blog-container { max-width: 980px; margin: 0 auto; margin-top: 20px;}
    textarea { height: 200px; }
</style>