package com.dev.test.main;


public class ObjectCreate {

    public static void main(String[] args){

        Object obj=new Object(){
            public boolean equals(Object obj){
                System.out.print(obj);
                return  true;
            }
        };


        System.out.println(obj.equals("fred"));

    }
}
