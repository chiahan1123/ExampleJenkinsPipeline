package com.example.someapp.utils;

import org.junit.Test;

import static org.junit.Assert.*;

public class StringsUnitTest {

    @Test
    public void testIsEmpty() throws Exception {
        assertTrue(Strings.isEmpty(null));
        assertTrue(Strings.isEmpty(""));
        assertTrue(Strings.isEmpty(" "));
    }

    @Test
    public void testIsNotEmpty() throws Exception {
        assertFalse(Strings.isEmpty("a"));
        assertFalse(Strings.isEmpty("ab c"));
    }

}
