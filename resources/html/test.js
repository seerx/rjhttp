Vue.component("card",{
    props:{//这里是组件可以接受的参数，也就是相当于面向原型写组件时的配置参数，用户可以传递不同参数，自己定义组件
        cardTitle:{//卡片标题
            type:String,
            default:'这是一个卡片'
        },
        list:{//列表内容
            type:Array,
            default:[]
        }
    },
    template:`
        <div class="modal">
            <div class="modal-header">
                <h4>{{cardTitle}}</h4>
            </div>
            <div class="modal-content">
                <div>
                    <slot name="modal-content">可以扩展的卡片内容</slot>
                    <ul>
                        <li v-for="(item,index) in list">{{item}}</li>
                    </ul>
                </div>
            </div>
            <div class="modal-footer">
                <slot name="modal-footer">
                    <input class="btn blue" type="button" value="ok" @click="okHandle" />
                    <input class="btn" type="button" value="cancel" @click="cancelHandle" />
                </slot> 
            </div>
        </div>
    `,
    methods:{//这里定义的组件的方法，利用$emit()进行父子组件通信，子组件通过点击事件告诉父组件触发一个自定义事件，$emit()方法第二个参数也可以用来传递数据
        okHandle(){
            this.$emit("ok");
        },
        cancelHandle(){
            this.$emit('cancel')
        }
    }
});
