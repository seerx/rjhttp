Vue.component("runner", {
    data () {
        return {
            rootUrl: '',
            title: "",
            root: null,
            rootIsArray: false,
            tree: null,
            json: '',
            response: '',
            complete: false,
            success: false
        }
    },
    mounted() {
    },
    methods: {
        init (rootUrl, serviceName, objectTree, rootObjectId, rootIsArray) {
            this.rootUrl = rootUrl
            this.title = serviceName
            this.tree = objectTree
            this.rootIsArray = rootIsArray
            this.root = objectTree[rootObjectId]
            this.genTemplate(rootObjectId)
        },
        genTemplate (rootId) {
            // 生成 json 参数
            let elementStack = []
            let arg = null
            if (rootId && rootId !== '') {
                arg = this.analyseObj(rootId, elementStack, this.rootIsArray)
            }
            let json = [{
                service: this.title,
                args: this.rootIsArray ? [arg] : arg
            }]
            this.json = JSON.stringify(json, null, '\t')
        },
        analyseObj (objId, stack, array) {
            const obj = this.tree[objId]
            let res = null
            const type = obj['type']
            switch (type) {
                case 'int':
                    res = 0
                    break
                case 'string':
                    res = ""
                    break
                case 'bool':
                    res = false
                    break
                default:
                    // 查找之前有没有
                    let found = false
                    res = null
                    for (var n = 0; n < stack.length; n ++) {
                        if (stack[n] === type) {
                            found = true
                            break
                        }
                    }
                    if (!found) {
                        res = {}
                        stack.push(type)
                        for (var n = 0; n < obj['children'].length; n ++) {
                            field = obj['children'][n]
                            res[field['name']] = this.analyseObj(field['reference'], stack, field['array'])
                        }
                        stack.pop()
                    }
            }
            if (array) {
                return [res]
            }
            return res
        },
        open() {
            window.open(this.rootUrl + '?' + this.json, '_blank')
        },
        get () {
            let ajax = new Ajax(this.rootUrl)
            ajax.Get(this.json).then(res => {
                this.complete = true
                let json = JSON.parse(res)
                this.success = !json['error']
                this.response = JSON.stringify(json, null, '  ')
                // console.log(res)
            }).catch(err => {
                this.complete = true
                this.success = false
                // console.log(err)
                this.response = err
            })
        },
        post () {
            let ajax = new Ajax(this.rootUrl)
            ajax.Post(this.json).then(res => {
                this.complete = true
                let json = JSON.parse(res)
                this.success = !json['error']
                this.response = JSON.stringify(json, null, '  ')
                // console.log(res)
            }).catch(err => {
                this.complete = true
                this.success = false
                // console.log(err)
                this.response = err
            })
        }
    },
    template: `<el-card class="box-card runner">
                    <div slot="header" class="clearfix">
                        <span>{{ title }}</span>
                        <el-button title="GET" v-if="json!==''" @click="get" style="margin-left: 6px; float: right; padding: 3px 3px" type="primary" icon="el-icon-bottom
"></el-button>    
                        <el-button title="POST" v-if="json!==''" @click="post" style="margin-left: 6px; float: right; padding: 3px 3px" type="primary" icon="el-icon-video-play
"></el-button>
                        <el-button title="OPEN" v-if="json!==''" @click="open" style="float: right; padding: 3px 3px" type="primary" icon="el-icon-download
"></el-button>
                    </div>
                    <el-input v-model="json" type="textarea" rows="10" style="width:100%;"></el-input> 
                    <div class="response" style="padding: 8px; border-bottom-radius: 4px; border: solid 1px rgb(220, 223, 230);">
                        <div v-if="complete && success">
                            <i class="el-icon-success" style="color: green;" type="success"></i> Success
                        </div>
                        <div v-if="complete && !success">
                            <i class="el-icon-error" style="color: red;" type="error"></i> Failed
                        </div>
                        <pre style="width: 100%;" >{{ response }}</pre>
                    </div>
                </el-card>`
})
