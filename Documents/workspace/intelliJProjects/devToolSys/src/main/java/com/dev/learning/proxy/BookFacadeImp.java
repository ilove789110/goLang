package com.dev.learning.proxy;

/**
 * Created by WilliamsLiu on 14/10/22.
 */

import static java.lang.System.out;

public class BookFacadeImp implements BookFacade {
    @Override
    public void read() {
        out.println("read books");
    }
}
