package com.dev.learning.proxy

import spock.lang.Specification

/**
 * Created by WilliamsLiu on 14/10/22.
 * 实现接口的类可以通过JDK的反射包进行动态代理
 */
class DynamicProxyByJDKTest extends Specification {



    public static main(String[] args){

        def dynamicProxy=new DynamicProxyByJDK()

        def bookFacade=(BookFacade) dynamicProxy.bind(new BookFacadeImp())

        bookFacade.read()

    }




}
