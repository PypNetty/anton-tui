package models

import "time"

// SystemMetrics représente les métriques système
type SystemMetrics struct {
	Timestamp time.Time
	CPU       CPUMetrics
	Memory    MemoryMetrics
	Disk      DiskMetrics
	Network   NetworkMetrics
}

// CPUMetrics représente les métriques CPU
type CPUMetrics struct {
	Usage       float64 // Pourcentage d'utilisation
	Temperature float64 // Température en Celsius
	Frequency   float64 // Fréquence en MHz
	Cores       []CoreMetrics
}

// CoreMetrics représente les métriques pour un cœur CPU
type CoreMetrics struct {
	ID          int
	Usage       float64
	Temperature float64
}

// MemoryMetrics représente les métriques de mémoire
type MemoryMetrics struct {
	Total        uint64
	Used         uint64
	Free         uint64
	UsagePercent float64
	SwapTotal    uint64
	SwapUsed     uint64
	SwapFree     uint64
}

// DiskMetrics représente les métriques de disque
type DiskMetrics struct {
	Total        uint64
	Used         uint64
	Free         uint64
	UsagePercent float64
	IORead       uint64
	IOWrite      uint64
	IOBusy       float64
}

// NetworkMetrics représente les métriques réseau
type NetworkMetrics struct {
	BytesSent     uint64
	BytesReceived uint64
	PacketsSent   uint64
	PacketsRecv   uint64
	Interfaces    []InterfaceMetrics
}

// InterfaceMetrics représente les métriques pour une interface réseau
type InterfaceMetrics struct {
	Name        string
	BytesSent   uint64
	BytesRecv   uint64
	PacketsSent uint64
	PacketsRecv uint64
	Errors      uint64
	Dropped     uint64
}
