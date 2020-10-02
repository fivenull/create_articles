<template>
  <div id="wrap">
    <Header/>
    <div class="container">
      <Message :show.sync="msgShow" :type="msgType" :msg="msg"/>
      <div class="row content-row">
        <Left/>
        <router-view v-if="isRouterAlive"/>
      </div>
    </div>
    <Footer/>
  </div>
</template>

<script>
  import Header from "@/components/layouts/Header"
  import Footer from "@/components/layouts/Footer"
  import Left from "@/components/pages/Left"
  export default {
    provide(){
      return{
        reload:this.reload
      }
    },
    components :{
      Header,
      Footer,
      Left
    },
    data(){
      return {
        msg: '', // 消息
        msgType: '', // 消息类型
        msgShow: false, // 是否显示消息，默认不显示
        isRouterAlive : true,
      }
    },
    methods: {
      showMsg(msg, type = 'success') {
        this.msg = msg
        this.msgType = type
        this.msgShow = true
      },
      reload(){
        this.isRouterAlive = false
        this.$nextTick(function(){
          this.isRouterAlive = true
        })
      }
    },
    mounted(){
      // console.log(this.$route.params);
    }
  }
</script>

<style lang="scss">
  $container-large-desktop: 1200px;
  $btn-primary-bg: #00b5ad;
  $btn-primary-border: #00b5ad;
  $label-primary-bg: #00b5ad;
  $pagination-active-bg: #00b5ad;
  $pagination-active-border: #00b5ad;
  $pagination-color: #00b5ad;
  $input-border-focus: #00b5ad;
  $link-color: #12c4c5;
  $link-hover-color: #22ddde;
  $icon-font-path: "~bootstrap-sass/assets/fonts/bootstrap/";
  $fa-font-path: "~font-awesome/fonts/";

  @import "~bootstrap-sass/assets/stylesheets/_bootstrap";
  @import "~font-awesome/scss/font-awesome";
  @import "./styles/main";
  @import "./styles/extra";
  @import '~simplemde/dist/simplemde.min.css';
  @import '~highlight.js/styles/paraiso-dark.css';
  .content-row{
    margin-bottom: 130px;
  }
</style>