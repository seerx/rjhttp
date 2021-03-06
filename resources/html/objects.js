Vue.component("objects", {
    data () {
        return {
            objectId: '',
            objectIsArray: false,
            objectIsRequire: false,
            objectMap: null,
            rootNode: null,
            resolve: null,
            expandArray: [],
            objectProps: {
                label: 'name',
                isLeaf: 'leaf'
            }
        }
    },
    methods: {
        init (objectId, objectIsArray, require, objectMap) {
            this.objectId = objectId
            this.objectIsRequire = require
            this.objectIsArray = objectIsArray
            this.objectMap = objectMap
            this.rootNode.childNodes = [];

            this.loadNode(this.rootNode, this.resolve)
            this.expandArray = [this.objectId]
        },
        generateDesc(data) {
            let desc = ''
            if (data['desc'] !== '') {
                desc = data['desc']
            }
            if (data['pattern']) {
                if (desc === '')  {
                    desc = 'pattern: ' + data['pattern']
                } else {
                    desc += '<br>pattern: ' + data['pattern']
                }
            }
            return desc
        },
        loadNode (node, resolve) {
            if (node.level === 0) {
                this.rootNode = node
                this.resolve = resolve
                // 根节点
                if (this.objectId === '' || !this.objectMap) {
                    return resolve([]);
                }
                obj = this.objectMap[this.objectId]
                // console.log(obj)
                if (obj) {
                    // this.expandArray = [obj['id']]
                    let name = this.objectIsArray ? '[' + obj['type'] + ']' : obj['type']
                    if (this.objectIsRequire) {
                        name = '*' + name
                    }
                    return resolve([{
                        id: obj['id'],
                        name: name,
                        info: obj,
                        require: this.objectIsRequire,
                        expanded: true,
                        leaf: obj['children'] === undefined
                    }]);
                } else {
                    return resolve([]);
                }
            }
            // console.log(node.data['info'])
            const data = []
            const items = node.data['info']['children']
            if (!items) {
                return resolve([]);
            }
            for (var n = 0; n < items.length; n ++) {
                const item = items[n]
                const obj = this.objectMap[item['reference']]
                let name = item['name'] + ':' + item['type']
                if (item['range']) {
                    name += ':' + item['range']
                }
                // if (item)
                if (item['array']) {
                    name = item['name'] + ':[' + item['type'] + ']'
                }
                if (item['require']) {
                    name = '*' + name
                }

                data.push({
                    id: obj['id'] + node.level,
                    name: name,
                    info: obj,
                    desc: item['description'],
                    deprecated: item['deprecated'],
                    range: item['range'],
                    pattern: item['pattern'],
                    require: item['require'],
                    // field: item,
                    // array: item['array'],
                    leaf: obj['children'] === undefined // ? true : false
                })
            }
            return resolve(data)
        }
    },

    template: `<el-tree :props="objectProps"
             :load="loadNode"
             :empty-text="'None'"
             :default-expanded-keys="expandArray"
             node-key="id"
             lazy>
        <span class="custom-tree-node" slot-scope="{ node, data }">
            <el-tooltip v-if="(data.desc && data.desc !== '') || data['pattern']" class="item" effect="dark" placement="right">
                <div slot="content">
                    <div v-if="data.desc !== ''">{{ data.desc }}</div>
                    <div v-if="data.pattern">pattern: {{ data.pattern }}</div>
                </div>
                <span :style="data['deprecated'] ? 'text-decoration: line-through;' : ''">{{ node.label }}</span>
            </el-tooltip>
            <span v-else :style="data['deprecated'] ? 'text-decoration: line-through;' : ''">{{ node.label }}</span>
            <span style="color: #aaa; margin-left: 16px; font-size: 14px;">
                {{ data.desc }}
            </span>

        </span>
    </el-tree>`,
    css: `body {
        background: #000;
    }`
})
