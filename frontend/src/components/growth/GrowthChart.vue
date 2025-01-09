<template>
  <div class="bg-white rounded-lg shadow p-6">
    <div class="flex justify-between items-center mb-4">
      <h3 class="text-lg font-semibold">{{ title }}</h3>
      <el-radio-group v-model="timeRange" size="small">
        <el-radio-button label="3">近3个月</el-radio-button>
        <el-radio-button label="6">近6个月</el-radio-button>
        <el-radio-button label="12">近1年</el-radio-button>
      </el-radio-group>
    </div>
    <div ref="chartRef" style="width: 100%; height: 400px"></div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import * as echarts from 'echarts'
import { formatDate } from '../../utils/date'

const props = defineProps({
  title: {
    type: String,
    required: true
  },
  records: {
    type: Array,
    required: true
  },
  unit: {
    type: String,
    required: true
  },
  standardData: {
    type: Object,
    default: () => ({})
  }
})

const chartRef = ref(null)
const chart = ref(null)
const timeRange = ref(3)

const initChart = () => {
  if (chart.value) {
    chart.value.dispose()
  }
  chart.value = echarts.init(chartRef.value)
}

const updateChart = () => {
  if (!chart.value || !props.records.length) return

  // 按日期排序
  const sortedRecords = [...props.records].sort((a, b) => new Date(a.date) - new Date(b.date))
  
  const dates = sortedRecords.map(r => formatDate(r.date, 'MM-DD'))
  const values = sortedRecords.map(r => r.value)

  const option = {
    tooltip: {
      trigger: 'axis',
      formatter: function(params) {
        const data = params[0]
        return `${data.name}<br/>${data.value}${props.unit}`
      }
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true
    },
    xAxis: {
      type: 'category',
      data: dates,
      boundaryGap: false
    },
    yAxis: {
      type: 'value',
      name: props.unit,
      nameLocation: 'end'
    },
    series: [
      {
        name: props.title,
        type: 'line',
        data: values,
        smooth: true,
        symbol: 'circle',
        symbolSize: 8,
        itemStyle: {
          color: '#6366f1'
        },
        lineStyle: {
          width: 3
        },
        areaStyle: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            {
              offset: 0,
              color: 'rgba(99, 102, 241, 0.3)'
            },
            {
              offset: 1,
              color: 'rgba(99, 102, 241, 0.1)'
            }
          ])
        }
      }
    ]
  }

  chart.value.setOption(option)
}

watch(() => props.records, updateChart, { deep: true })
watch(timeRange, updateChart)

onMounted(() => {
  initChart()
  updateChart()
  
  window.addEventListener('resize', () => {
    chart.value?.resize()
  })
})
</script> 