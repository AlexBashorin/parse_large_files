package parse

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const ipv4Size = 1 << 32

type IPv4Counter struct {
	bitArray    []uint64
	uniqueCount uint32
	uniq_ip     []string
}

func NewIPv4Counter() *IPv4Counter {
	return &IPv4Counter{
		bitArray: make([]uint64, ipv4Size),
	}
}

func (c *IPv4Counter) AddIP(ip string) {
	ipInt := ipToInt(ip)
	index, bit := ipInt/64, ipInt%64
	mask := uint64(1) << bit

	if c.bitArray[index]&mask == 0 {
		c.bitArray[index] |= mask
		c.uniqueCount++
		c.uniq_ip = append(c.uniq_ip, ip)
	}
}

func ipToInt(ip string) uint32 {
	octets := strings.Split(ip, ".")
	var result uint32
	for i := 0; i < 4; i++ {
		octet, _ := strconv.Atoi(octets[i])
		result = result<<8 | uint32(octet)
	}
	return result
}

func CountUniqueIPs(filename string) (uint32, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	counter := NewIPv4Counter()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		counter.AddIP(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	fmt.Println(counter.uniq_ip)
	if len(counter.uniq_ip) > 0 {
		for i := 0; i < len(counter.uniq_ip); i++ {
			unic_ips, err := os.OpenFile("./unique.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

			if err != nil {
				return 0, err
			}
			defer unic_ips.Close()

			unic_ips.WriteString(counter.uniq_ip[i] + "\n")
		}
	}

	return counter.uniqueCount, nil
}
