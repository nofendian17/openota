// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: proto/country/v1/country_service.proto

// Protobuf Java Version: 3.25.1
package com.core.country.v1;

public interface GetAllResponseOrBuilder extends
    // @@protoc_insertion_point(interface_extends:core.country.v1.GetAllResponse)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>repeated .core.country.v1.Country countries = 1 [json_name = "countries"];</code>
   */
  java.util.List<com.core.country.v1.Country> 
      getCountriesList();
  /**
   * <code>repeated .core.country.v1.Country countries = 1 [json_name = "countries"];</code>
   */
  com.core.country.v1.Country getCountries(int index);
  /**
   * <code>repeated .core.country.v1.Country countries = 1 [json_name = "countries"];</code>
   */
  int getCountriesCount();
  /**
   * <code>repeated .core.country.v1.Country countries = 1 [json_name = "countries"];</code>
   */
  java.util.List<? extends com.core.country.v1.CountryOrBuilder> 
      getCountriesOrBuilderList();
  /**
   * <code>repeated .core.country.v1.Country countries = 1 [json_name = "countries"];</code>
   */
  com.core.country.v1.CountryOrBuilder getCountriesOrBuilder(
      int index);
}