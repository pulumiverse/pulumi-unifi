// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.unifi.iam.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetUserResult {
    /**
     * @return Specifies whether this user should be blocked from the network.
     * 
     */
    private Boolean blocked;
    /**
     * @return Override the device fingerprint.
     * 
     */
    private Integer devIdOverride;
    /**
     * @return fixed IPv4 address set for this user.
     * 
     */
    private String fixedIp;
    /**
     * @return The hostname of the user.
     * 
     */
    private String hostname;
    /**
     * @return The ID of the user.
     * 
     */
    private String id;
    /**
     * @return The IP address of the user.
     * 
     */
    private String ip;
    /**
     * @return The local DNS record for this user.
     * 
     */
    private String localDnsRecord;
    /**
     * @return The MAC address of the user.
     * 
     */
    private String mac;
    /**
     * @return The name of the user.
     * 
     */
    private String name;
    /**
     * @return The network ID for this user.
     * 
     */
    private String networkId;
    /**
     * @return A note with additional information for the user.
     * 
     */
    private String note;
    /**
     * @return The name of the site the user is associated with.
     * 
     */
    private String site;
    /**
     * @return The user group ID for the user.
     * 
     */
    private String userGroupId;

    private GetUserResult() {}
    /**
     * @return Specifies whether this user should be blocked from the network.
     * 
     */
    public Boolean blocked() {
        return this.blocked;
    }
    /**
     * @return Override the device fingerprint.
     * 
     */
    public Integer devIdOverride() {
        return this.devIdOverride;
    }
    /**
     * @return fixed IPv4 address set for this user.
     * 
     */
    public String fixedIp() {
        return this.fixedIp;
    }
    /**
     * @return The hostname of the user.
     * 
     */
    public String hostname() {
        return this.hostname;
    }
    /**
     * @return The ID of the user.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The IP address of the user.
     * 
     */
    public String ip() {
        return this.ip;
    }
    /**
     * @return The local DNS record for this user.
     * 
     */
    public String localDnsRecord() {
        return this.localDnsRecord;
    }
    /**
     * @return The MAC address of the user.
     * 
     */
    public String mac() {
        return this.mac;
    }
    /**
     * @return The name of the user.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The network ID for this user.
     * 
     */
    public String networkId() {
        return this.networkId;
    }
    /**
     * @return A note with additional information for the user.
     * 
     */
    public String note() {
        return this.note;
    }
    /**
     * @return The name of the site the user is associated with.
     * 
     */
    public String site() {
        return this.site;
    }
    /**
     * @return The user group ID for the user.
     * 
     */
    public String userGroupId() {
        return this.userGroupId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetUserResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean blocked;
        private Integer devIdOverride;
        private String fixedIp;
        private String hostname;
        private String id;
        private String ip;
        private String localDnsRecord;
        private String mac;
        private String name;
        private String networkId;
        private String note;
        private String site;
        private String userGroupId;
        public Builder() {}
        public Builder(GetUserResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.blocked = defaults.blocked;
    	      this.devIdOverride = defaults.devIdOverride;
    	      this.fixedIp = defaults.fixedIp;
    	      this.hostname = defaults.hostname;
    	      this.id = defaults.id;
    	      this.ip = defaults.ip;
    	      this.localDnsRecord = defaults.localDnsRecord;
    	      this.mac = defaults.mac;
    	      this.name = defaults.name;
    	      this.networkId = defaults.networkId;
    	      this.note = defaults.note;
    	      this.site = defaults.site;
    	      this.userGroupId = defaults.userGroupId;
        }

        @CustomType.Setter
        public Builder blocked(Boolean blocked) {
            this.blocked = Objects.requireNonNull(blocked);
            return this;
        }
        @CustomType.Setter
        public Builder devIdOverride(Integer devIdOverride) {
            this.devIdOverride = Objects.requireNonNull(devIdOverride);
            return this;
        }
        @CustomType.Setter
        public Builder fixedIp(String fixedIp) {
            this.fixedIp = Objects.requireNonNull(fixedIp);
            return this;
        }
        @CustomType.Setter
        public Builder hostname(String hostname) {
            this.hostname = Objects.requireNonNull(hostname);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder ip(String ip) {
            this.ip = Objects.requireNonNull(ip);
            return this;
        }
        @CustomType.Setter
        public Builder localDnsRecord(String localDnsRecord) {
            this.localDnsRecord = Objects.requireNonNull(localDnsRecord);
            return this;
        }
        @CustomType.Setter
        public Builder mac(String mac) {
            this.mac = Objects.requireNonNull(mac);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder networkId(String networkId) {
            this.networkId = Objects.requireNonNull(networkId);
            return this;
        }
        @CustomType.Setter
        public Builder note(String note) {
            this.note = Objects.requireNonNull(note);
            return this;
        }
        @CustomType.Setter
        public Builder site(String site) {
            this.site = Objects.requireNonNull(site);
            return this;
        }
        @CustomType.Setter
        public Builder userGroupId(String userGroupId) {
            this.userGroupId = Objects.requireNonNull(userGroupId);
            return this;
        }
        public GetUserResult build() {
            final var o = new GetUserResult();
            o.blocked = blocked;
            o.devIdOverride = devIdOverride;
            o.fixedIp = fixedIp;
            o.hostname = hostname;
            o.id = id;
            o.ip = ip;
            o.localDnsRecord = localDnsRecord;
            o.mac = mac;
            o.name = name;
            o.networkId = networkId;
            o.note = note;
            o.site = site;
            o.userGroupId = userGroupId;
            return o;
        }
    }
}
