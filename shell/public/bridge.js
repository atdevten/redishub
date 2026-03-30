(function () {
  /**
   * Scorix Bridge Orchestrator
   * This file decides whether to use AppBridge (WebView) or WebBridge (WebSocket)
   */
  const scorix = {
    _bridge: null,
    _initPromise: null,

    async init(options = {}) {
      if (this._initPromise) return this._initPromise;

      this._initPromise = (async () => {
        // 1. Detect environment
        if (window.__scorix__ipc_emit) {
          console.debug("Scorix: App Mode detected (WebView)");
          this._bridge = window.ScorixAppBridge;
        } else {
          console.debug("Scorix: Web Mode detected (WebSocket)");
          this._bridge = window.ScorixWebBridge;
        }

        // 2. Initialize the chosen bridge
        if (this._bridge && typeof this._bridge.init === 'function') {
          await this._bridge.init(options);
        }
      })();

      return this._initPromise;
    },

    async invoke(method, params, options) {
      await this.init();
      if (!this._bridge) throw new Error("Scorix: Bridge failed to initialize.");
      return this._bridge.invoke(method, params, options);
    },

    async emit(topic, data) {
      await this.init();
      if (!this._bridge) throw new Error("Scorix: Bridge failed to initialize.");
      return this._bridge.emit(topic, data);
    },

    async on(topic, callback) {
      await this.init();
      if (!this._bridge) throw new Error("Scorix: Bridge failed to initialize.");
      return this._bridge.on(topic, callback);
    },

    async resolve(name, handler) {
      await this.init();
      if (!this._bridge) throw new Error("Scorix: Bridge failed to initialize.");
      return this._bridge.resolve(name, handler);
    }
  };

  if (typeof window !== "undefined") {
    window.scorix = scorix;
    window.scorix.init().catch(console.error);
  }
})();
