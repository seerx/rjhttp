Vue.component("runner", {
    data () {
        return {
            title: "",
            root: null,
            rootIsArray: false,
            tree: null,
            json: ''
        }
    },
    mounted() {
    },
    methods: {
        init (serviceName, objectTree, rootObjectId, rootIsArray) {
            this.title = serviceName
            this.tree = objectTree
            this.rootIsArray = rootIsArray
            this.root = objectTree[rootObjectId]
            console.log(this.root)
            this.genTemplate(rootObjectId)
        },
        genTemplate (rootId) {
            // 生成 json 参数
            let elementStack = []
            let arg = this.analyseObj(rootId, elementStack, this.rootIsArray)
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
                    res = {}
                    for (var n = 0; n < stack.length; n ++) {
                        if (stack[n] === type) {
                            found = true
                            break
                        }
                    }
                    if (!found) {
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
        run () {

        }
    },
    template: `<el-card class="box-card runner">
                    <div slot="header" class="clearfix">
                        <span>{{ title }}</span>
                        <el-button @click="run" style="float: right; padding: 3px 3px" type="primary" icon="el-icon-video-play
"></el-button>
                    </div>
                    <el-input v-model="json" type="textarea" rows="10" style="width:100%;"></el-input> 
                    <div class="response"> 
                        <pre style="width: 100%;"></pre>
                    </div>
                </el-card>`
})
