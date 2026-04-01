"use client"

export interface ScorixAPI {
  invoke<T = any>(method: string, params?: any, options?: any): Promise<T>
  emit(topic: string, data?: any): Promise<void>
  on(topic: string, callback: (data: any, error: string) => void): () => void
  resolve(name: string, handler: (data: any) => any): void
  init(options?: any): Promise<void>
}

declare global {
  interface Window {
    scorix?: ScorixAPI
    __scorix__ipc_emit?: (msg: string) => Promise<any>
    ScorixWebBridge?: {
      _status: "connected" | "connecting" | "disconnected"
    }
  }
}

const scorix: ScorixAPI = {
  invoke: (...args) => {
    if (typeof window === "undefined" || !window.scorix) {
      throw new Error("Scorix API not available")
    }
    return window.scorix.invoke(...args)
  },
  emit: (...args) => {
    if (typeof window === "undefined" || !window.scorix) return Promise.resolve()
    return window.scorix.emit(...args)
  },
  on: (...args) => {
    if (typeof window === "undefined" || !window.scorix) return () => {}
    return window.scorix.on(...args)
  },
  resolve: (...args) => {
    if (typeof window === "undefined" || !window.scorix) return
    return window.scorix.resolve(...args)
  },
  init: (...args) => {
    if (typeof window === "undefined" || !window.scorix) return Promise.resolve()
    return window.scorix.init(...args)
  },
}

export default scorix
