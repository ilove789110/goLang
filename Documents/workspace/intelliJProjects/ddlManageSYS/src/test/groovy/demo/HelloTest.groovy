package demo

import spock.lang.Specification

/**
 * Created by WilliamsLiu on 14-9-13.
 */
class HelloTest extends Specification {
    def "SayHello"() {
        expect:
        hello == "Hello,Gradle!"

        where:
        hello = new Hello().sayHello()
    }
}
