<script setup lang="ts">
import { reactive, defineProps, onMounted } from 'vue';
import PulseLoader from 'vue-spinner/src/PulseLoader.vue';
import { type Job } from '@/models/Jobs'
import JobListings from '@/components/JobListings.vue';
import axios from 'axios';

type JobsResponse = {
    code: number;
    data: {
        jobs: Job[];
    };
}

defineProps({
    limit: Number,
    showAll: {
        type: Boolean,
        default: false
    },
});

const state = reactive({
    jobs: <Job[]>([]),
    isLoading: true,
});

onMounted(async () => {
    try {
        const response = await axios.get<JobsResponse>('http://localhost:8080/jobs');
        console.log('Jobs fetched:', response.data.data.jobs);
        state.jobs = response.data.data.jobs;
    } catch (error) {
        console.error('Error fetching jobs:', error);
    } finally {
        state.isLoading = false;
    }
})

</script>

<template>
    <section class="bg-blue-50 px-4 py-10">
        <div class="container-xl lg:container m-auto">
            <h2 class="text-3xl font-bold text-blue-500 mb-6 text-center">
                Browse Jobs
            </h2>
            <!-- Show Loading Spinner while reading data -->
            <div v-if="state.isLoading" class="text-center text-gray-500 py-6">
                <PulseLoader />
            </div>
            <!-- Show Jobs -->
            <div v-else class="grid grid-cols-1 md:grid-cols-3 gap-6">
                <JobListings v-for="job in state.jobs.slice(0, limit || state.jobs.length)" :key="job.id" :job="job" />
            </div>
        </div>
    </section>

    <section v-if="showAll" class="m-auto max-w-lg my-10 px-6">
      <a
        href="/jobs"
        class="block bg-black text-white text-center py-4 px-6 rounded-xl hover:bg-gray-700"
        >View All Jobs</a
      >
    </section>
    
</template>