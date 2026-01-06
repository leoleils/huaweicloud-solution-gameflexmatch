// Copyright (c) Huawei Technologies Co., Ltd. 2022-2022. All rights reserved.

// Package cloudprovider 云厂商抽象层 - 工厂模式
package cloudprovider

import (
	"fmt"
	"sync"
)

var (
	providersMu sync.RWMutex
	providers   = make(map[string]func() CloudProvider)
)

// RegisterProvider 注册云厂商
func RegisterProvider(name string, factory func() CloudProvider) {
	providersMu.Lock()
	defer providersMu.Unlock()
	providers[name] = factory
}

// NewProvider 创建云厂商实例
func NewProvider(config *ProviderConfig) (CloudProvider, error) {
	providersMu.RLock()
	factory, ok := providers[config.ProviderName]
	providersMu.RUnlock()

	if !ok {
		return nil, fmt.Errorf("unknown cloud provider: %s", config.ProviderName)
	}

	provider := factory()
	if err := provider.Initialize(config); err != nil {
		return nil, fmt.Errorf("failed to initialize provider %s: %v", config.ProviderName, err)
	}

	return provider, nil
}

// providerCache 云厂商实例缓存
var providerCache = struct {
	sync.RWMutex
	cache map[string]CloudProvider
}{cache: make(map[string]CloudProvider)}

// GetProvider 获取已注册的云厂商(带缓存)
func GetProvider(config *ProviderConfig) (CloudProvider, error) {
	cacheKey := fmt.Sprintf("%s_%s_%s", config.ProviderName, config.Region, config.ProjectId)

	providerCache.RLock()
	if p, ok := providerCache.cache[cacheKey]; ok {
		providerCache.RUnlock()
		return p, nil
	}
	providerCache.RUnlock()

	providerCache.Lock()
	defer providerCache.Unlock()

	// Double check
	if p, ok := providerCache.cache[cacheKey]; ok {
		return p, nil
	}

	p, err := NewProvider(config)
	if err != nil {
		return nil, err
	}

	providerCache.cache[cacheKey] = p
	return p, nil
}

// GetProviderWithCredentials 使用凭证获取云厂商实例
func GetProviderWithCredentials(providerName, region, projectId, ak, sk string, endpoints map[string]string) (CloudProvider, error) {
	config := &ProviderConfig{
		ProviderName: providerName,
		Region:       region,
		ProjectId:    projectId,
		AccessKey:    ak,
		SecretKey:    sk,
		Endpoints:    endpoints,
	}
	return GetProvider(config)
}

// ClearProviderCache 清除云厂商缓存
func ClearProviderCache() {
	providerCache.Lock()
	defer providerCache.Unlock()
	providerCache.cache = make(map[string]CloudProvider)
}

// RemoveProviderFromCache 从缓存中移除指定的云厂商实例
func RemoveProviderFromCache(providerName, region, projectId string) {
	cacheKey := fmt.Sprintf("%s_%s_%s", providerName, region, projectId)
	providerCache.Lock()
	defer providerCache.Unlock()
	delete(providerCache.cache, cacheKey)
}

// ListRegisteredProviders 列出已注册的云厂商
func ListRegisteredProviders() []string {
	providersMu.RLock()
	defer providersMu.RUnlock()

	names := make([]string, 0, len(providers))
	for name := range providers {
		names = append(names, name)
	}
	return names
}
