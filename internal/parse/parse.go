package parse

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Определяем размер битового массива
const ipv4Size = 1 << 32

// The structure includes: a bitmap, a counter for the number of unique values, and an array of unique values.
// структура включает: битой массив, счетчик количества уникальных значений и массив уникальных значений
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

// adding a bit to our bit array
// добавление бита в нашем битовом массиве
func (c *IPv4Counter) AddIP(ip string) {
	// converting to 32-bit arr
	ipInt := ipToInt(ip)
	// get index for array and bit
	index, bit := ipInt/64, ipInt%64
	// create mask 64-bit and place our bit
	mask := uint64(1) << bit

	// AND between 64-bit array and mask: if array don't have value in this index, place our bit (with mask)
	// 10101100 and 00010000 - we will got 0 (1 & 0 = 0)
	// побитовая операция AND между значением из массива и текущей маской, если нет занчения (0), добавляем
	// например, 10101100 and 00010000 - получим 0, т.к. 1 & 0 = 0
	if c.bitArray[index]&mask == 0 {
		// OR with same: got 1 (1 | 0 = 1) and assigning a value (bit to the desired index)
		// присваиваем плюс побитовое ИЛИ: если значение установлено - не трогаем, иначе записываем
		c.bitArray[index] |= mask
		c.uniqueCount++
		// append this unique IP
		// добавляем в массив (тут тудушка: сделать проверку на уникальность по ходу, а не просто записывать в отдельный массив)
		c.uniq_ip = append(c.uniq_ip, ip)
	}
	// TODO:
	// remove not unique IP in the process, without additional array
}

// converting a string representation of an IP to a 32-bit number
// преобразование строкового представления IP в 32 битное число
func ipToInt(ip string) uint32 {
	octets := strings.Split(ip, ".")
	var result uint32
	for i := 0; i < 4; i++ {
		octet, _ := strconv.Atoi(octets[i])
		result = result<<8 | uint32(octet)
	}
	return result
}

func CountUniqueIPs(filename string) (uint32, string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, "", err
	}
	defer file.Close()

	counter := NewIPv4Counter()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		counter.AddIP(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return 0, "", err
	}

	if len(counter.uniq_ip) > 0 {
		// TODO: Error: open ./unique.txt: too many open files
		for i := 0; i < len(counter.uniq_ip); i++ {
			unic_ips, err := os.OpenFile("./unique.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

			if err != nil {
				return 0, "", err
			}
			defer unic_ips.Close()

			unic_ips.WriteString(counter.uniq_ip[i] + "\n")
		}
	}

	return counter.uniqueCount, "./unique.txt", nil
}
