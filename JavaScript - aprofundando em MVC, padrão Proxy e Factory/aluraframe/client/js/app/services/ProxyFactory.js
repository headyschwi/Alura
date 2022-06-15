class ProxyFactory{

    static create(model, props, action){
        return new Proxy(model, {
            
            get(target, prop, receiver) {
                
                if(props.includes(prop) && typeof(target[prop]) == typeof(Function)) {
                    
                    return function() {
                        
                        console.log(`interceptando ${prop}`);
                        const retorno = Reflect.apply(target[prop], target, arguments);
                        action(target);
                        return retorno
                    }
                }
                
                return Reflect.get(target, prop, receiver);
            },

            set(target, prop, value, receiver){
                const retorno = Reflect.set(target, prop, value, receiver)
                
                if(props.includes(prop)){
                    action(target)
                }
                return retorno
            }

            
        });
    }
}