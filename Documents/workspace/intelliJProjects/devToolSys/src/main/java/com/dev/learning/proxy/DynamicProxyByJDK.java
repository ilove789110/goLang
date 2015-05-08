package com.dev.learning.proxy;

import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;
import java.lang.reflect.Proxy;

import static java.lang.System.out;

/**
 * Created by WilliamsLiu on 14/10/22.
 */
public class DynamicProxyByJDK implements InvocationHandler {

    Object target;

    public Object bind(Object target){

        this.target=target;

        return Proxy.newProxyInstance(target.getClass().getClassLoader(),target.getClass().getInterfaces(),this);

    }


    @Override
    public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
        Object result=null;

        out.println("before method");

        result=method.invoke(this.target,args);

        out.println("after method");

        return result;
    }
}
