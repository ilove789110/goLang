package com.dev.learning.proxy

import spock.lang.Specification

/**
 * Created by WilliamsLiu on 14/10/22.
 * 没有实现接口的类需要使用Cglib来进行动态代理
 */
class BookFacadeCglibTest extends Specification {

    public static void main(String[] args)
    {
        def cglib=new BookFacadeCglib()

        def bookImp=(BookImp) cglib.getInstance(new BookImp())

        bookImp.read()
    }

}
