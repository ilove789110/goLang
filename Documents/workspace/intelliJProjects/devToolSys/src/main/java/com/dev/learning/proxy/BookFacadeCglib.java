package com.dev.learning.proxy;


import net.sf.cglib.proxy.Enhancer;
import net.sf.cglib.proxy.MethodInterceptor;
import net.sf.cglib.proxy.MethodProxy;

import java.lang.reflect.Method;

import static java.lang.System.out;

/**
 * Created by WilliamsLiu on 14/10/22.
 */
public class BookFacadeCglib implements MethodInterceptor {

    private Object target;

    public Object getInstance(Object target)
    {
        this.target=target;

        Enhancer enhancer=new Enhancer();

        enhancer.setSuperclass(target.getClass());

        enhancer.setCallback(this);

        return enhancer.create();

    }

    @Override
    public Object intercept(Object o, Method method, Object[] args, MethodProxy methodProxy) throws Throwable {

        Object result=null;

        out.println("before method");

        result=methodProxy.invokeSuper(o,args);

        out.println("after method");

        return result;

    }
}
