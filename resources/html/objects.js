Vue.component("objects", {
    data () {
        return {
            objectId: '',
            objectIsArray: false,
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
        init (objectId, objectIsArray, objectMap) {
            this.objectId = objectId
            this.objectIsArray = objectIsArray
            this.objectMap = objectMap
            this.rootNode.childNodes = [];

            this.loadNode(this.rootNode, this.resolve)
            this.expandArray = [this.objectId]
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
                    return resolve([{
                        id: obj['id'],
                        name: this.objectIsArray ? '[' + obj['type'] + ']' : obj['type'],
                        info: obj,
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
                if (item['array']) {
                    name = item['name'] + ':[' + item['type'] + ']'
                }
                data.push({
                    id: obj['id'] + node.level,
                    name: name,
                    info: obj,
                    desc: item['description'],
                    deprecated: item['deprecated'],
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
            <el-tooltip v-if="data.desc && data.desc !== ''" class="item" effect="dark" :content="data.desc" placement="right">
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
