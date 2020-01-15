package pages

const IndexContext = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Run JSON</title>
    <link rel="stylesheet" href="https://unpkg.com/element-ui@2.13.0/lib/theme-chalk/index.css">
</head>
<body>
    <div id="app">
        <el-container>
            <el-header>
                <a class="logo" href="https://github.com/seerx/runjson">
                    <img src="https://raw.githubusercontent.com/seerx/runjson/master/resources/logo.png">
                    <span>Run JSON</span>
                </a>
            </el-header>
            <el-container>
                <el-aside width="260px" class="list">
                    <el-collapse v-model="activeName" accordion>
                        <el-collapse-item v-for="(grp, n) in data['groups']" :title="grp.name + grp.description" :name="n" :key="n">
                            <div class="service" v-for="(svc, n) in grp['services']" :key="n">
                                <a @click="showInfo(svc)"><h4>{{ svc.name }}</h4></a>
                                <span>{{ svc.description }}</span>
                            </div>
                        </el-collapse-item>
                    </el-collapse>
                </el-aside>

                <el-aside width="400px" class="info">
                    <el-card class="box-card">
                        <div slot="header" class="clearfix">
                            <span>{{ current ? current.name : 'Please select service'}}</span>
                            <el-button v-if="current" @click="initParam" style="float: right; padding: 3px 3px" type="primary" icon="el-icon-arrow-right"></el-button>
                        </div>
                        <el-card v-if="current" class="desc">
                            {{ current.description }}
                        </el-card>
                        <div class="object">
                            <h4>Request</h4>
                            <objects ref="request"></objects>
                        </div>
                        <div class="object">
                            <h4>Response</h4>
                            <objects ref="response"></objects>
                        </div>
                    </el-card>
                </el-aside>
                <el-main>
                    <runner ref="runner" :rootUrl="rootUrl">
                    </runner>
                </el-main>
            </el-container>
        </el-container>
    </div>
</body>
<!-- import Vue before Element -->
<script src="https://unpkg.com/vue@v2.6.11/dist/vue.min.js"></script>
<!-- import JavaScript -->
<script src="https://unpkg.com/element-ui@2.13.0/lib/index.js"></script>

<script src="?m=file&file=objects.js"></script>
<script src="?m=file&file=runner.js"></script>
<script src="?m=file&file=ajax.js"></script>

<script>
    new Vue({
        el: '#app',
        data: function() {
            return {
                rootUrl: '',
                visible: false,
                activeName: 0,
                current: null,
                list:[1,2,3,4],
                data: [],
                request: {
                    rootNode: null,
                    resolve: null
                },
                response: {
                    rootNode: null,
                    resolve: null
                },
                objectProps: {
                    // children: 'children',
                    label: 'name',
                    isLeaf: 'leaf'
                }
            }
        },
        mounted () {
            loc = document.location
            this.rootUrl = loc.pathname
            let ajax = new Ajax(this.rootUrl)
            ajax.Get('m=graph').then(res => {
                this.data = JSON.parse(res)
            }).catch(err => {
                console.log(res)
            })
        },
        methods: {
            initParam () {
                let isAry = this.current['inputIsArray']
                let objId = this.current['inputObjectId']
                let tree = this.data['request']
                this.$refs.runner.init(this.rootUrl, this.current['name'],  tree, objId, isAry)
            },
            showInfo (info) {
                this.current = info
                let reqiure = this.current['inputIsRequire']
                this.$refs.request.init(info['inputObjectId'],
                    info['inputIsArray'],
                    reqiure,
                    this.data['request'])
                this.$refs.response.init(info['outputObjectId'],
                    info['outputIsArray'],
                    false,
                    this.data['response'])
            }
        }
    })
</script>

<style>
    html,body {
        margin: 0;
        padding: 0;
        height: 100%;
    }

    body #app {
        height: 100%;
    }

    .el-main {
        background-color: #E9EEF3;
        color: #333;
        padding: 2px;
    }

    body > #app > .el-container {
        height: 100%;
        margin-bottom: 0px;
    }

    .logo {
        height: 36px;
        cursor: pointer;
        display: flex;
        flex-flow: row;
        align-items: flex-end;
    }

    .logo img, .logo span {
        height: 100%;
        line-height: 36px;
        width: auto;
        vertical-align: middle;
    }

    .el-header, .el-footer {
        background: linear-gradient(#f7f7f7, #e2e2e2);
        background: -webkit-gradient(linear, left top, left bottom, from(#f7f7f7), to(#e2e2e2));
        color: #333;
        height: 36px !important;
        margin: 0;
        padding: 0;
    }


    .el-aside {
        background-color: #D3DCE6;
        color: #333;
        text-align: center;
        padding: 2px;
    }

    .list .el-collapse-item__header {
        padding: 0 0 0 8px !important;
    }

    .list .el-collapse-item__content {
        padding: 0 !important;
        text-align: left;
    }


    .service {
        padding: 0 0 8px 25px !important;
    }

    .service h4 {
        margin: 0;
        cursor: pointer;
    }



    /*.el-container:nth-child(5) .el-aside,*/
    /*.el-container:nth-child(6) .el-aside {*/
    /*    line-height: 100%;*/
    /*}*/

    /*.el-container:nth-child(7) .el-aside {*/
    /*    line-height: 100%;*/
    /*}*/

    .info {
    }

    .info>.el-card {
        height: 100%;
        background-color: #f7f7f7;
    }

    .info .el-divider {
        margin: 10px 0 8px 0 !important;
    }

    .info .el-card__body {
        padding: 8px;
        overflow: auto;
        height: calc(100% - 79px)
    }

    .info .el-card__header {
        text-align: left;
        padding: 8px 12px 8px 12px;
    }

    .info .object {
        /*width: 100%;*/
    }

    .info .object h4 {
        margin: 6px 0 6px 0;
    }

    .info .desc {
        margin-bottom: 8px;
        text-align: left;
    }

    .custom-tree-node {
        flex: 1;
        display: flex;
        align-items: center;
        justify-content: space-between;
        font-size: 17px;
        padding-right: 8px;
    }

    .clearfix:before,
    .clearfix:after {
        display: table;
        content: "";
    }
    .clearfix:after {
        clear: both
    }


    .runner {
        height: 100%;
        overflow: auto;
    }
    .runner .el-card__header {
        text-align: left;
        padding: 8px 12px 8px 12px;
    }
    .runner .el-card__body {
        height: calc(100% - 55px) !important;
        padding: 8px;
    }
    .runner .el-card__body .el-textarea__inner {
        resize: none;
    }

    .runner .el-card__header .el-input {
        width: 180px !important;
        margin-top: 0px;
        margin-bottom: -1px;
    }
    .runner .el-card__header .el-input .el-input__inner {
        padding-bottom: 0px;
        line-height: 22px;
        height: 22px;
        padding-left: 3px;
        padding-right: 3px;
        border-bottom-right-radius: 0px;
        border-top-right-radius: 0px;
    }
    .runner .el-card__header .btn {
        margin-left: 6px;
        float: right;
        padding: 3px 3px;
    }
    .runner .el-card__header .upload {
        margin-left: 0px;
        border-left: none;
        border-bottom-left-radius: 0px;
        border-top-left-radius: 0px;
    }
</style>

</html>
`