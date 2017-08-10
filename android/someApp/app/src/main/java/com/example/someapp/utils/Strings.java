package com.example.someapp.utils;

import android.support.annotation.Nullable;

/**
 * Strings utility class.
 */
public final class Strings {

    /**
     * Determines whether the target {@link String} is empty. Both a {@code null} string and an
     * empty string will return {@code true}.
     *
     * @param string target {@link String}.
     * @return a boolean to indicate whether is target {@link String} is empty.
     */
    public static boolean isEmpty(@Nullable String string) {
        return (string == null || string.trim().length() == 0);
    }

    /**
     * Capitalize the first letter of the target {@link String}.  If the target {@link String}
     * is {@code null}, then return {@code null}.  If the target {@link String} is empty, then
     * return an empty {@link String}. Otherwise, return a {@link String} with a capitalized first
     * letter.
     *
     * @param string target {@link String}.
     * @return a capitalized {@link String}.
     */
    @Nullable public static String capitalize(@Nullable String string) {
        if (string == null) {
            return null;
        }
        if (isEmpty(string)) {
            return "";
        }
        final char firstChar = string.charAt(0);
        if (Character.isUpperCase(firstChar)) {
            return string;
        } else {
            return Character.toUpperCase(firstChar) + string.substring(1);
        }
    }

    private Strings() {}

}
