// Code generated by protoapi; DO NOT EDIT.

package com.yoozoo.spring;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.List;

public class SearchKeyValueListRequest {
    private final String key;
    private final int service_id;
    private final int env_id;

    @JsonCreator
    public SearchKeyValueListRequest(@JsonProperty("key") String key, @JsonProperty("service_id") int service_id, @JsonProperty("env_id") int env_id) {
        this.key = key;
        this.service_id = service_id;
        this.env_id = env_id;
    }

    public String getKey() {
        return key;
    }
    public int getService_id() {
        return service_id;
    }
    public int getEnv_id() {
        return env_id;
    }
    
}