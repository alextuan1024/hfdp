# 工厂模式
## 披萨店

- 将`PizzaStore`作为虚基类
- `createPizza()`方法设置为abstract，由派生类实现
- `orderPizza()`方法调用派生类的`createPizza()`方法

## Golang实现披萨店的问题

- 由于没有基类概念，golang中如果使用`interface`代替虚基类，那么prepare/cut/box等方法需要重复实现
- `interface`有别于`abstract Class`，无法包含`field`和`method`的默认实现
- 可选的策略模式实现，好处是可以动态修改`Prepare(),Bake(),Cut(),Box()`行为
    
    ```go
    type Preparer interface{
        Prepare()
    }
    type Pizza interface{
        PerformPrepare()
    }
    type NYStyleCheesePizza struct{
        name        string
        preparer    Preparer
    }
    func (ny *NYStyleCheesePizza)PerformPrepare(){
        // 默认的prepare实现
        if ny.preparer == nil{
            fmt.Printf("Preparing %s\n",Name)
        }
        ny.preparer.Prepare()
    }
    // 实现Preparer接口，特定的Prepare()方法
    type NYStyleCheesePizzaPrepare struct{}
    func (ncp *NYStyleCheesePizzaPrepare)Prepare(){
        //
    }
    ```
    