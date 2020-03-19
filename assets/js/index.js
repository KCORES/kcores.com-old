new Vue({
    el: '#content',
    data() {
        return {
            info: null,
            serverErrored: false,
            originalURL: "", // user inpputed origin url for generate token.
            urlInput: {
                color: "#0A0A0F",
            },
            generateType: "0", // token generate type
            responseError: false,
            redirectLoop: false,
            illegalURI: false,
        }
    },
    filters: {
        currencydecimal(value) {
            return value.toFixed(2)
        }
    },
    mounted() {

    },
    watch: {

    },
    methods: {
        generateToken() {
            axios
                .get("http://" + window.location.host + "/generate?url=" + this.originalURL + "&GenerateType=" + this.generateType + "&RedirectType=5")
                .then(response => {
                    if (response.data.s != 1) {
                        if (response.data.m == "redirect loop") {
                            this.redirectLoop = true
                        } else if (response.data.m == "illegal URI") {
                            this.illegalURI = true
                        }
                        this.originalURL = response.data.m
                        this.urlInput.color = "#e83929"
                        this.responseError = true

                    } else {
                        this.info = response.data.d.ID
                        this.originalURL = window.location.host + "/" + response.data.d.ID
                        this.urlInput.color = "#00a497"
                    }
                })
                .catch(error => {
                    this.serverErrored = true
                })
                .finally()
        },
        // clean color for next input
        cleanColor() {
            this.urlInput.color = "#0A0A0F"
        },
        // when server response error and user click to reinput, clean error info
        cleanErrorInfo() {
            if (this.responseError) {
                this.responseError = false
                this.redirectLoop = false
                this.illegalURI = false
                this.originalURL = ""
                this.urlInput.color = "#0A0A0F"
            }
        }
    }
})