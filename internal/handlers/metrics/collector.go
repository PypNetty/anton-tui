package metrics

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"

	"anton-tui/internal/models"
)

// Collector gère la collecte des métriques système
type Collector struct {
	interval time.Duration
	metrics  chan models.SystemMetrics
}

// NewCollector crée un nouveau collecteur
func NewCollector(interval time.Duration) *Collector {
	return &Collector{
		interval: interval,
		metrics:  make(chan models.SystemMetrics, 100),
	}
}

// Start démarre la collecte des métriques
func (c *Collector) Start() {
	go func() {
		ticker := time.NewTicker(c.interval)
		defer ticker.Stop()

		for range ticker.C {
			metrics, err := c.collect()
			if err != nil {
				continue
			}
			c.metrics <- metrics
		}
	}()
}

// collect récupère toutes les métriques système
func (c *Collector) collect() (models.SystemMetrics, error) {
	metrics := models.SystemMetrics{
		Timestamp: time.Now(),
	}

	// CPU
	cpuPercent, err := cpu.Percent(0, false)
	if err == nil && len(cpuPercent) > 0 {
		metrics.CPU.Usage = cpuPercent[0]
	}

	// Mémoire
	if vmstat, err := mem.VirtualMemory(); err == nil {
		metrics.Memory = models.MemoryMetrics{
			Total:        vmstat.Total,
			Used:         vmstat.Used,
			Free:         vmstat.Free,
			UsagePercent: vmstat.UsedPercent,
		}
	}

	// Disque
	if diskstat, err := disk.Usage("/"); err == nil {
		metrics.Disk = models.DiskMetrics{
			Total:        diskstat.Total,
			Used:         diskstat.Used,
			Free:         diskstat.Free,
			UsagePercent: diskstat.UsedPercent,
		}
	}

	// Réseau
	if netstat, err := net.IOCounters(false); err == nil && len(netstat) > 0 {
		metrics.Network = models.NetworkMetrics{
			BytesSent:     netstat[0].BytesSent,
			BytesReceived: netstat[0].BytesRecv,
			PacketsSent:   netstat[0].PacketsSent,
			PacketsRecv:   netstat[0].PacketsRecv,
		}
	}

	return metrics, nil
}

// GetMetrics retourne le canal des métriques
func (c *Collector) GetMetrics() <-chan models.SystemMetrics {
	return c.metrics
}

// Stop arrête la collecte des métriques
func (c *Collector) Stop() {
	close(c.metrics)
}
