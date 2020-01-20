Vue.component("runner", {
    data () {
        return {
            rootUrl: '',
            title: '',
            tokenField: '',
            tokenValue: '',
            root: null,
            rootIsArray: false,
            tree: null,
            json: '',
            response: '',
            complete: false,
            success: false,
            file: 'file',
            requestCount: 0
        }
    },
    mounted() {
    },
    methods: {
        setToken(field, value) {
            this.tokenField = field
            this.tokenValue = value
        },
        init (rootUrl, serviceName, objectTree, rootObjectId, rootIsArray) {
            this.rootUrl = rootUrl
            this.title = serviceName
            this.tree = objectTree
            this.rootIsArray = rootIsArray
            this.root = objectTree[rootObjectId]
            this.genTemplate(rootObjectId)
            this.response = ''
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
                    res = ''
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
        headers() {
            if (! this.tokenField || !this.tokenValue) {
                return null
            }
            if (this.tokenField === '') {
                return null
            }
            let h = {}
            h[this.tokenField] = this.tokenValue
            return h
        },
        upload () {
            if (this.file === '') {
                this.$message({
                    message: 'Please input upload file\'s field name',
                    type: 'error'
                });
                return
            }
            this.$refs.file.click()
        },
        doUpload () {
            this.requestCount ++
            const file = this.$refs.file.files[0];
            const ajax = new Ajax(this.rootUrl, this.headers())
            ajax.Upload(this.json, file, this.file).then(res => {
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
        open() {
            // this.requestCount ++
            window.open(this.rootUrl + '?' + this.json, '_blank')
        },
        get () {
            this.requestCount ++
            let ajax = new Ajax(this.rootUrl, this.headers())
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
            this.requestCount ++
            let ajax = new Ajax(this.rootUrl, this.headers())
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
    template: `<el-card class='box-card runner'>
                    <div slot='header' class='clearfix'>
                        <span>{{ title }}</span>
                        <el-button title='GET' v-if="json!==''" @click="get" class='btn' type='success' icon='el-icon-bottom
'></el-button>    
                        <el-button title='POST' v-if="json!==''" @click="post" class='btn' type='primary' icon='el-icon-video-play
'></el-button>
                        
                        <el-button title='OPEN' v-if="json!==''" @click="open" class='btn' type='info' icon='el-icon-link
'></el-button>
                        <el-button title='UPLOAD' v-if="json!==''" @click="upload" class='upload btn' type='warning' icon='el-icon-upload
'></el-button><input ref='file' type='file' style='display: none;' @change="doUpload">
                        <el-input v-if="json!==''" placeholder='Please input upload field' style='float: right;'
                            v-model="file">
                        </el-input>
                    </div>
                    <el-input v-model="json" type='textarea' rows='10' style='width:100%;'></el-input> 
                    <div class='response' style='padding: 8px; border-bottom-radius: 4px; border: solid 1px rgb(220, 223, 230);'>
                        <div v-if="complete && success">
                            <i class='el-icon-success' style='color: rgb(103, 194, 58);' type='success'></i> Success [{{requestCount}}]
                        </div>
                        <div v-if="complete && !success">
                            <i class='el-icon-error' style='color: red;' type='error'></i> Failed [{{requestCount}}]
                        </div>
                        
                        <pre style='width: 100%;' >{{ response }}</pre>
                    </div>
                </el-card>`
})
